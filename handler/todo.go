package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pirosiki197/todo_list/repository"
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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := validateTodo(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id, err := h.repo.CreateTodo(toTodoModel(todo))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	todo.ID = id

	return c.JSON(http.StatusCreated, todo)
}

// GET /todos/{id}
func (h *Handler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	todoModel, err := h.repo.GetTodo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	todo := Todo{
		ID:     todoModel.ID,
		Task:   todoModel.Task,
		Status: todoModel.Status,
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

	todoModel, err := h.repo.GetTodo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	todo := Todo{
		ID:     todoModel.ID,
		Task:   todoModel.Task,
		Status: req.Status,
	}
	if err := validateTodo(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = h.repo.UpdateTodo(id, toTodoModel(todo))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func toTodoModel(todo Todo) repository.Todo {
	return repository.Todo{
		ID:     todo.ID,
		Task:   todo.Task,
		Status: todo.Status,
	}
}
