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
	content := templates.Index("something", pages.BodyContent())
	// Define template body content.
	e.GET("/", func(c echo.Context) error {
		return cmd.Render(c, 200, content)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
