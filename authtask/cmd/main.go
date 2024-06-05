package main

import (
	"github.com/novychok/go-samples/authtask/internal/pkg/psql"
	prodRepo "github.com/novychok/go-samples/authtask/internal/repository/product"
	signRepo "github.com/novychok/go-samples/authtask/internal/repository/sign"

	l "github.com/novychok/go-samples/authtask/internal/pkg/log"

	"github.com/novychok/go-samples/authtask/internal/handler"

	prodSrv "github.com/novychok/go-samples/authtask/internal/service/product"
	signSrv "github.com/novychok/go-samples/authtask/internal/service/sign"

	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	port := flag.String("port", "5456", "server port")
	flag.Parse()

	l := l.New()

	db, err := psql.New()
	if err != nil {
		l.Error("Failed to connect to psql", "err", err.Error())
	}
	if err := initTables(db, l); err != nil {
		l.Error("Failed to init psql tables", "err", err.Error())
	}

	signRepo := signRepo.New(db)
	prodRepo := prodRepo.New(db)

	signService := signSrv.New(signRepo, l)
	prodService := prodSrv.New(prodRepo, l)

	handler := handler.New(signService, prodService)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: handler.InitRoutes(),
	}

	server := New(httpServer)

	l.Error(fmt.Sprintf("Failed to run server on port: %s", *port),
		server.httpServer.ListenAndServe())
}

func initTables(db *sql.DB, l *slog.Logger) error {
	usersTable := `
    CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(250) NOT NULL
	);`

	_, err := db.Exec(usersTable)
	if err != nil {
		return fmt.Errorf("error to create users table")
	}

	productTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		product_name VARCHAR(255) UNIQUE NOT NULL,
		product_description TEXT NOT NULL,
		product_price VARCHAR(255) NOT NULL
	);`

	_, err = db.Exec(productTable)
	if err != nil {
		return fmt.Errorf("error to create products table")
	}

	insertIntoProducts := `
	INSERT INTO products (id, product_name, product_description, product_price)
	VALUES
    ('1', 'Product1', 'Description for Product 1', '10.99'),
    ('2', 'Product2', 'Description for Product 2', '20.49'),
    ('3', 'Product3', 'Description for Product 3', '15.99'),
    ('4', 'Product4', 'Description for Product 4', '5.99'),
    ('5', 'Product5', 'Description for Product 5', '8.99');`

	_, err = db.Exec(insertIntoProducts)
	if err != nil {
		return fmt.Errorf("error to insert in products table")
	}

	l.Info("Successfully created all psql tables")
	return nil
}

func dropTables(db *sql.DB, l *slog.Logger) error {

	dropProducts := `DELETE * FROM products;`

	_, err := db.Exec(dropProducts)
	if err != nil {
		return fmt.Errorf("error to delete * from products")
	}

	l.Info("Successfully dropped all products tables")

	return nil
}

func New(httpServer *http.Server) *Server {
	return &Server{
		httpServer: httpServer,
	}
}
