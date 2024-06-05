package handler

import (
	"net/http"
	"os"

	"github.com/novychok/go-samples/authtask/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignMiddleware(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, "user is not authenticated")
	}

	var registeredClaims entity.RegisteredClaims
	if err := h.signService.ParseAccessToken(token, &registeredClaims, os.Getenv("TOKEN_SECRET")); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, "user is not authenticated")
	}

	c.Next()
}
