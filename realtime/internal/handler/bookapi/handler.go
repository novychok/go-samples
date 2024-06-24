package bookapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/novychok/go-samples/realtime/internal/entity"
	"github.com/novychok/go-samples/realtime/internal/service"
)

type Handler struct {
	bookapiService service.Books
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {

	var bookRequest entity.Book
	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		http.Error(w, fmt.Sprintf("incorrect request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	bookId, err := h.bookapiService.Create(r.Context(), &bookRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't create the book: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", bookId)))
}

func New(bookapiService service.Books) *Handler {
	return &Handler{
		bookapiService: bookapiService,
	}
}
