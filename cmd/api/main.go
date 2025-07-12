package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ymatsukawa/backbase-go/cmd/plugin"
	"github.com/ymatsukawa/backbase-go/pkg/env"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db := plugin.NewDatabase()
	conn, err := db.Launch()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	layers := plugin.NewLayers(conn, logger)
	layers.Build()
	h := layers.GetHandlers()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(plugin.CustomCors))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", h.UserHandler.GetAllUsers)

	port := env.GetEnv("SERVER_PORT", "8080")
	addr := fmt.Sprintf(":%s", port)

	e.Start(addr)
}
