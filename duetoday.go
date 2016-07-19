package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func main() {
	e := echo.New()

	e.GET("/", index)

	log.Println("listening...")
	e.Run(standard.New(":4001"))
}
