package interceptor

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
)

func Unarylogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Trace("got request : ", info.FullMethod, ":", req)
	resp, err = handler(ctx, req)
	log.Trace("served resp : ", info.FullMethod, ":", resp, ":", err)
	return resp, err
}
