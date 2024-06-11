package handler

import (
	"fmt"
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

func (todo *Todo) SetStatus(status string) error {
	new := *todo
	new.Status = status
	if err := validateTodo(new); err != nil {
		return err
	}
	todo.Status = status
	return nil
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func validateTodo(todo Todo) error {
	switch todo.Status {
	case StatusProcessing, StatusDone:
		return nil
	default:
		return fmt.Errorf("unknown todo status: %s", todo.Status)
	}
}

// POST /todos
func (h *Handler) CreateTodo(c echo.Context) error {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := validateTodo(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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

// PATCH /todos/{id}
func (h *Handler) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var req UpdateStatusRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	todo, ok := h.todos[id]
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}
	if err := todo.SetStatus(req.Status); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	h.todos[id] = todo

	return c.NoContent(http.StatusOK)
}
