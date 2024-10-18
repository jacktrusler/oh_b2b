package main

import (
	"htmx-fr/cmd"
	"htmx-fr/templates"
	"htmx-fr/templates/pages"

	"github.com/labstack/echo/v4"
)

func main() {
	// Start Echo server
	e := echo.New()
	e.Static("/public", "public")
	// Define template body content.
	e.GET("/", func(c echo.Context) error {
		content := templates.Index("something", pages.Home())
		return cmd.Render(c, 200, content)
	})
	e.GET("/lmao", func(c echo.Context) error {
		content := templates.Index("something", pages.Lmao())
		return cmd.Render(c, 200, content)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
