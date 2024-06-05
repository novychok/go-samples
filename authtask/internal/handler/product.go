package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProduct(c *gin.Context) error {
	productName := c.Param("productName")

	product, err := h.productService.GetByName(c.Request.Context(), productName)
	if err != nil {
		return errorResponse{
			Message: fmt.Sprintf("error while get product by name: %s", err.Error()),
			Status:  http.StatusBadRequest,
		}
	}

	writeStatusResponse(c, http.StatusOK, product)
	return nil
}
