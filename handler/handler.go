package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pirosiki197/todo_list/repository"
)

type Handler struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) SetUpRoutes(e *echo.Echo) {
	e.GET("/ping", h.Ping)

	e.POST("/todos", h.CreateTodo)
	e.GET("/todos/:id", h.GetTodo)
	e.PATCH("/todos/:id", h.UpdateStatus)
}

func (h *Handler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
