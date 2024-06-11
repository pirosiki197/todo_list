package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetUpRoutes(e *echo.Echo) {
	e.GET("/ping", h.Ping)
}

func (h *Handler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
