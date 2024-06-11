package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	todos map[int]Todo
}

func New() *Handler {
	return &Handler{
		todos: make(map[int]Todo),
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
