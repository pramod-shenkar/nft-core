package server

import (
	"context"
	"net"
	"nftledger/common/config"
	"nftledger/internal/adapter/inbound/grpc/interceptor"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/grpc"

	"go.uber.org/fx"
)

type App struct {
	*fiber.App
	*grpc.Server
}

func NewApp() *App {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.Unarylogger,
		),
	)
	return &App{app, server}
}

func Run(lc fx.Lifecycle, a *App) error {

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				listerner, err := net.Listen("tcp", ":"+config.App.Service.Port)
				if err != nil {
					return err
				}

				if config.App.Service.EndPointType == config.Rest {
					go func() {
						err = a.Listener(listerner)
						if err != nil {
							return
						}
					}()
				} else {
					err = a.Serve(listerner)
					if err != nil {
						return err
					}
				}

				log.Info("nft-core service started at : ")

				return nil
			},

			OnStop: func(ctx context.Context) error {
				return a.Shutdown()
			},
		},
	)

	return nil
}
