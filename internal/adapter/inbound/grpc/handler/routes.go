package grpc

import (
	"encoding/json"
	"nft-protoc/generated/service"
	"nftledger/common/server"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc/reflection"
)

func RegisterRoute(a *server.App, handlers *Handlers) {
	service.RegisterRequestServiceServer(a.Server, handlers.Request)
	reflection.Register(a.Server)

	payload, _ := json.MarshalIndent(a.GetServiceInfo(), " ", " ")
	log.Info(string(payload))
}
