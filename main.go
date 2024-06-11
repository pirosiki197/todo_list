package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pirosiki197/todo_list/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	handler := handler.New()
	handler.SetUpRoutes(e)

	e.Start(":8080")
}
