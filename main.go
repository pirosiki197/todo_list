package main

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pirosiki197/todo_list/handler"
	"github.com/pirosiki197/todo_list/repository"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	db := setupDB()
	repo := repository.New(db)

	handler := handler.New(repo)
	handler.SetUpRoutes(e)

	e.Start(":8080")
}

func setupDB() *sql.DB {
	config := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "todo_list",
		Loc:       time.Local,
		ParseTime: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
