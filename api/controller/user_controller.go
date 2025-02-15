package controller

import (
	"api/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetById(ctx echo.Context) error
}

type userController struct {
	userService service.UserService
}

// GetById implements UserController.
func (u *userController) GetById(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := u.userService.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}
