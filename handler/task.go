package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(200, "tasks")
	}
}

func PutTasks(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		return ctx.JSON(200, H{
			"deleted": id,
		})
	}
}
