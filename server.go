package main

import (
	"net/http"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create renderer
	gv := echoview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/application",
	})

	// Renderer
	e.Renderer = gv

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{
		"title": "Hello World!",
	})
}
