package main

import (
	"hello"
	"logger"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/test", func(c echo.Context) error {
		logger.Info("Sending hello world")
		return c.JSON(http.StatusOK, hello.GetHelloWorld())
		//return c.JSON(http.StatusOK, "Hello")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
