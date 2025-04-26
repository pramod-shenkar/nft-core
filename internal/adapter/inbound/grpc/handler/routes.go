package grpc

import (
	"nftledger/common/server"
	"nftledger/internal/adapter/inbound/grpc/generated/service"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc/reflection"
)

func RegisterRoute(a *server.App, handlers *Handlers) {
	service.RegisterRequestServiceServer(a.Server, handlers.Request)
	reflection.Register(a.Server)
	log.Infof("%+v", a.GetServiceInfo())
}
