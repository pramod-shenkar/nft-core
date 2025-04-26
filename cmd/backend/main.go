package main

import (
	"context"
	"nftledger/common/server"
	grpc "nftledger/internal/adapter/inbound/grpc/handler"
	"nftledger/internal/adapter/inbound/rest"
	"nftledger/internal/adapter/outbound/dao"
	"nftledger/internal/adapter/outbound/dao/db"
	"nftledger/internal/core/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(

		fx.Provide(

			db.New,
			dao.NewRequestDao,
			dao.NewDaos,
			api.NewRequestService,
			api.NewServices,
			rest.NewRequestHandler,
			rest.NewHandlers,
			grpc.NewRequestHandler,
			grpc.NewHandlers,
			fiber.New,
			server.NewApp,
		),
		fx.Invoke(rest.RegisterRoute),
		fx.Invoke(grpc.RegisterRoute),

		fx.Invoke(server.Run),
	)
	err := app.Start(context.Background())
	if err != nil {
		log.Error(err.Error())
		return
	}
}
