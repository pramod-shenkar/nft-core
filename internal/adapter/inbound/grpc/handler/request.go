package grpc

import (
	"context"
	"fmt"
	. "nft-protoc/generated/model"
	. "nft-protoc/generated/service"
	"nftledger/internal/adapter/inbound/grpc/bind"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"nftledger/internal/core/api"
	"nftledger/internal/core/model"
)

type RequestHandler struct {
	*api.Services
	UnimplementedRequestServiceServer
}

func NewRequestHandler(
	api *api.Services,
) *RequestHandler {
	return &RequestHandler{
		api,
		UnimplementedRequestServiceServer{},
	}
}

func (h *RequestHandler) SaveRequest(c context.Context, req *SaveRequestRequest) (*SaveRequestResponse, error) {

	var request *model.Request = &model.Request{}

	if req.Request != nil {
		if req.Request.Name != "" {
			request.Name = req.Request.Name
		}
		if req.Request.Filetype != "" {
			request.Filetype = req.Request.Filetype
		}
		if req.Request.Filecontent != "" {
			request.FileContent = []byte(req.Request.Filecontent)
		}

	}

	err := h.RequestService.SaveRequest(request)
	if err != nil {
		return &SaveRequestResponse{Error: err.Error()}, err
	}

	return &SaveRequestResponse{Status: "sucess"}, nil
}

func (h *RequestHandler) GetRequest(c context.Context, req *GetRequestRequest) (*GetRequestResponse, error) {

	var where *model.Request = &model.Request{}
	if req.Where != nil {
		if req.Where.Id != 0 {
			where.ID = uint(req.Where.Id)
		}
		if req.Where.Name != "" {
			where.Name = req.Where.Name
		}
		if req.Where.CreatedAt != nil {
			where.CreatedAt = req.Where.CreatedAt.AsTime()
		}
		if req.Where.UpdatedAt != nil {
			where.UpdatedAt = req.Where.UpdatedAt.AsTime()
		}
	}

	request, err := h.RequestService.GetRequest(where)
	if err != nil {
		return nil, err
	}

	response := &GetRequestResponse{
		Request: &Request{
			Id:        uint64(request.ID),
			Name:      request.Name,
			CreatedAt: timestamppb.New(request.CreatedAt),
			UpdatedAt: timestamppb.New(request.UpdatedAt),
		},
	}

	return response, nil
}

func (h *RequestHandler) GetRequests(req *GetRequestsRequest, stream grpc.ServerStreamingServer[GetRequestsResponse]) error {

	var where *model.Request = &model.Request{}
	if req.Where != nil {
		if req.Where.Id != 0 {
			where.ID = uint(req.Where.Id)
		}
		if req.Where.Name != "" {
			where.Name = req.Where.Name
		}
		if req.Where.CreatedAt != nil {
			where.CreatedAt = req.Where.CreatedAt.AsTime()
		}
		if req.Where.UpdatedAt != nil {
			where.UpdatedAt = req.Where.UpdatedAt.AsTime()
		}
	}

	requests, err := h.RequestService.GetRequests(where)
	if err != nil {
		return err
	}

	go func() error {
		for {
			select {
			case <-stream.Context().Done():
				return fmt.Errorf("client cancelled stream")
			case request := <-model.RequestNotifier:
				stream.Send(&GetRequestsResponse{Request: bind.ModelToGrpcRequest(request)})
			}
		}
	}()

	for _, request := range *requests {
		model.RequestNotifier <- &request
	}

	return nil
}

func (h *RequestHandler) UpdateRequest(c context.Context, req *UpdateRequestRequest) (*UpdateRequestResponse, error) {

	var where *model.Request = &model.Request{}
	if req.Where != nil {
		if req.Where.Id != 0 {
			where.ID = uint(req.Where.Id)
		}
		if req.Where.Name != "" {
			where.Name = req.Where.Name
		}
		if req.Where.Status != 0 {
			where.Status.FromInt(int32(req.Where.Status))
		}
		if req.Where.Filetype != "" {
			where.Filetype = req.Where.Filetype
		}

		if req.Where.CreatedAt != nil {
			where.CreatedAt = req.Where.CreatedAt.AsTime()
		}
		if req.Where.UpdatedAt != nil {
			where.UpdatedAt = req.Where.UpdatedAt.AsTime()
		}
	}

	var update *model.Request = &model.Request{}
	if req.Request != nil {
		if req.Request.Id != 0 {
			update.ID = uint(req.Request.Id)
		}
		if req.Request.Name != "" {
			update.Name = req.Request.Name
		}
		if req.Request.Status != 0 {
			update.Status.FromInt(int32(req.Request.Status))
		}
		if req.Request.Filetype != "" {
			update.Filetype = req.Request.Filetype
		}
		if req.Request.CreatedAt != nil {
			update.CreatedAt = req.Request.CreatedAt.AsTime()
		}
		if req.Request.UpdatedAt != nil {
			update.UpdatedAt = req.Request.UpdatedAt.AsTime()
		}
	}

	err := h.RequestService.UpdateRequest(where, update)
	if err != nil {
		return nil, err
	}

	return &UpdateRequestResponse{Status: "success"}, nil
}

func (h *RequestHandler) DeleteRequest(c context.Context, req *DeleteRequestRequest) (*DeleteRequestResponse, error) {

	var where *model.Request = &model.Request{}
	if req.Where != nil {
		if req.Where.Id != 0 {
			where.ID = uint(req.Where.Id)
		}
		if req.Where.Name != "" {
			where.Name = req.Where.Name
		}
		if req.Where.CreatedAt != nil {
			where.CreatedAt = req.Where.CreatedAt.AsTime()
		}
		if req.Where.UpdatedAt != nil {
			where.UpdatedAt = req.Where.UpdatedAt.AsTime()
		}
	}

	err := h.RequestService.DeleteRequest(where)
	if err != nil {
		return nil, err
	}

	return &DeleteRequestResponse{Status: "sucess"}, nil
}
