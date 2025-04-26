package bind

import (
	. "nftledger/internal/adapter/inbound/grpc/generated/model"
	"nftledger/internal/core/model"
)

func ModelToGrpcRequest(*model.Request) *Request {
	return &Request{}
}

func GrpcToModelRequest(*Request) *model.Request {
	return &model.Request{}
}
