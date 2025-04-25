package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go.uber.org/fx"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	return &App{app}
}

func Run(lc fx.Lifecycle, a *App) error {

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return a.Listen(":9999")
			},

			OnStop: func(ctx context.Context) error {
				return a.Shutdown()
			},
		},
	)

	return nil
}
