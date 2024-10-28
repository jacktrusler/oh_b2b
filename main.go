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
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// StoreUser stores the username and hashed password in the SQLite DB
func StoreUser(db *sql.DB, username, hashedPassword string) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, username, hashedPassword)
	return err
}

func main() {
	// Start Echo server
	e := echo.New()
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// statement, _ := db.Prepare("DROP TABLE IF EXISTS users")
	// statement.Exec()
	//
	statement, _ := db.Prepare(
		`CREATE TABLE IF NOT EXISTS users 
    (
      id INTEGER PRIMARY KEY, 
      username TEXT UNIQUE, 
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
	e.POST("/signup", func(c echo.Context) error {
		usr := c.FormValue("username")
		pass := c.FormValue("password")

		hashedPassword, err := HashPassword(pass)
		if err != nil {
			log.Println("Error hashing password", err)
			return err
		}
		err = StoreUser(db, usr, hashedPassword)
		if err != nil {
			if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
				log.Println("Username already exists:", err)
				content := components.ErrorMsg("Username already exists", "Sign Up")
				return cmd.Render(c, 409, content)
			}
			log.Println("Error storing user:", err)
			content := components.ErrorMsg(err.Error(), "Sign Up")
			return cmd.Render(c, 500, content)
		}
		fmt.Println("User registered successfully!")
		content := components.SuccessMsg("You've been successfully registered!")
		return cmd.Render(c, 200, content)
	})

	e.POST("/login", func(c echo.Context) error {
		usr := c.FormValue("username")
		// pass := c.FormValue("password")
		var err error = nil
		var statement *sql.Stmt
		query := "SELECT * FROM users WHERE users.username = ?"
		statement, err = db.Prepare(query)
		rows, err := statement.Query(usr)

		if err != nil {
			log.Fatal(err)
		}

		// Check if a row was returned
		if rows.Next() {
			fmt.Println("User found:", usr)
			content := components.SuccessMsg("You've been successfully logged in!")
			return cmd.Render(c, 200, content)
		} else {
			fmt.Println("User not found")
		}
		return err
	})

	e.Logger.Fatal(e.Start(":1323"))
}
