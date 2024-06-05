package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/novychok/go-samples/authtask/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) error {
	var signUpRequest entity.SignUp

	if err := c.BindJSON(&signUpRequest); err != nil {
		return errorResponse{
			Message: "invalid input data",
			Status:  http.StatusBadRequest,
		}
	}

	generatedJwtToken, err := h.signService.SignUp(c.Request.Context(), &signUpRequest)
	if err != nil {
		return errorResponse{
			Message: fmt.Sprintf("error while sign up the user: %s", err.Error()),
			Status:  http.StatusBadRequest,
		}
	}

	setCookie(c, generatedJwtToken)
	writeStatusResponse(c, http.StatusOK, "signed up")
	return nil
}

func (h *Handler) SignIn(c *gin.Context) error {
	var signInRequest entity.SignIn

	if err := c.BindJSON(&signInRequest); err != nil {
		return errorResponse{
			Message: "invalid input data",
			Status:  http.StatusBadRequest,
		}
	}

	generatedJwtToken, err := h.signService.SignIn(c.Request.Context(), &signInRequest)
	if err != nil {
		return errorResponse{
			Message: fmt.Sprintf("error while sign in the user: %s", err.Error()),
			Status:  http.StatusBadRequest,
		}
	}

	setCookie(c, generatedJwtToken)
	writeStatusResponse(c, http.StatusOK, "logged in")
	return nil
}

func setCookie(c *gin.Context, generatedJwtToken string) {
	cookieTime, err := strconv.Atoi(os.Getenv("COOKIE_EXPIRATION"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid cookie time"))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", generatedJwtToken,
		cookieTime, "", "", false, true)
}
