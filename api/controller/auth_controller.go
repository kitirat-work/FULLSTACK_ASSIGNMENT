package controller

import (
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginPin(ctx echo.Context) error
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// LoginPin implements AuthController.
type LoginPinRequest struct {
	ID  string `json:"id"`
	Pin string `json:"pin"`
}

func (a *authController) LoginPin(ctx echo.Context) error {
	req := new(LoginPinRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user, err := a.authService.LoginPin(ctx.Request().Context(), req.ID, req.Pin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}
