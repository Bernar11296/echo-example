package main

import (
	"database/sql"
	"fmt"
	"log"

	"echo-check/handler"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDb("storage.db")
	migrate(db)
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handler.GetTasks(db))
	e.PUT("/tasks", handler.PutTasks(db))
	e.DELETE("/tasks", handler.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8080"))
}

func initDb(fp string) *sql.DB {
	db, err := sql.Open("sqlite3", fp)
	if err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}
	if db == nil {
		fmt.Println(2)
		log.Fatal("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	fmt.Println(3)
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(4)
		log.Fatal(err)
	}
}
