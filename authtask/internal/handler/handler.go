package handler

import (
	"github.com/novychok/go-samples/authtask/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	signService    service.SignService
	productService service.ProductService
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", handleWithError(h.SignUp))
	router.POST("/sign-in", handleWithError(h.SignIn))

	router.GET("/api/v1/products/:productName", h.SignMiddleware,
		handleWithError(h.GetProduct))

	return router
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e errorResponse) Error() string {
	return e.Message
}

type handleErrorFunc func(c *gin.Context) error

func handleWithError(h handleErrorFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			if e, ok := err.(errorResponse); ok {
				writeStatusResponse(c, e.Status, e.Message)
			}
		}
	}
}

func writeStatusResponse(c *gin.Context, code int, a any) {
	c.JSON(code, a)
}

func New(signService service.SignService, productService service.ProductService) *Handler {
	return &Handler{
		signService:    signService,
		productService: productService,
	}
}
