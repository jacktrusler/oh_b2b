package main

import (
	"database/sql"
	"fmt"
	"htmx-fr/cmd"
	"htmx-fr/templates"
	"htmx-fr/templates/components"
	"htmx-fr/templates/pages"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Start Echo server
	e := echo.New()
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, _ := db.Prepare("DROP TABLE IF EXISTS users")
	statement.Exec()

	statement, _ = db.Prepare(
		`CREATE TABLE IF NOT EXISTS users 
    (
      id INTEGER PRIMARY KEY, 
      username TEXT, 
      password TEXT, 
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
	statement.Exec()

	e.Static("/public", "public")
	// Define template body content.
	e.GET("/", func(c echo.Context) error {
		content := templates.Index("Ohio BSD", pages.Home(), components.Header(), components.Footer())
		return cmd.Render(c, 200, content)
	})
	e.GET("/lmao", func(c echo.Context) error {
		content := templates.Index("something", pages.Lmao(), components.Header(), components.Footer())
		return cmd.Render(c, 200, content)
	})
	e.POST("/login", func(c echo.Context) error {
		usr := c.FormValue("username")
		pass := c.FormValue("password")
		var err error = nil
		var statement *sql.Stmt
		statement, err = db.Prepare(`INSERT INTO users (username,password) VALUES (?, ?)`)
		_, err = statement.Exec(usr, pass)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(statement)

		return err
	})

	e.Logger.Fatal(e.Start(":1323"))
}
