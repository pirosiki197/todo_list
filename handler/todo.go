package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

const (
	StatusProcessing = "processing"
	StatusDone       = "done"
)

// POST /todos
func (h *Handler) CreateTodo(c echo.Context) error {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// 本当は排他制御が必要
	if _, ok := h.todos[todo.ID]; ok {
		return c.NoContent(http.StatusConflict)
	}

	todo.Status = StatusProcessing
	h.todos[todo.ID] = todo

	return c.NoContent(http.StatusCreated)
}

// GET /todos/{id}
func (h *Handler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	todo, ok := h.todos[id]
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, todo)
}
