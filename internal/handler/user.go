package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ymatsukawa/backbase-go/internal/response"
	"github.com/ymatsukawa/backbase-go/internal/service"
)

type UserHandler interface {
	GetAllUsers(ctx echo.Context) error
}

type userHandler struct {
	userService service.UserService
	Logger      *slog.Logger
}

func NewUserHandler(userService service.UserService, logger *slog.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		Logger:      logger,
	}
}

func (c *userHandler) GetAllUsers(ctx echo.Context) error {
	users, err := c.userService.GetAllUsers(ctx.Request().Context())
	if err != nil {
		slog.ErrorContext(ctx.Request().Context(), "failed to get all users", "server error", err)
		ctx.JSON(http.StatusInternalServerError, response.SimpleError500)
	}

	r := response.ToUsersResponse(users)
	return ctx.JSON(http.StatusOK, r)
}
