package grpc

import (
	"context"
	. "nftledger/internal/adapter/inbound/grpc/generated/model"
	. "nftledger/internal/adapter/inbound/grpc/generated/service"

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
		request.Name = req.Request.Name
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

func (h *RequestHandler) GetRequests(c context.Context, req *GetRequestsRequest) (*GetRequestsResponse, error) {

	var request *model.Request = &model.Request{}
	if req.Where != nil {
		request.Name = req.Where.Name
		request.CreatedAt = req.Where.CreatedAt.AsTime()
		request.UpdatedAt = req.Where.UpdatedAt.AsTime()
	}

	requests, err := h.RequestService.GetRequests(request)
	if err != nil {
		return nil, err
	}

	var response []*Request
	for _, request := range *requests {
		response = append(response, &Request{
			Id:        uint64(request.ID),
			Name:      request.Name,
			CreatedAt: timestamppb.New(request.CreatedAt),
			UpdatedAt: timestamppb.New(request.UpdatedAt),
		})
	}

	return &GetRequestsResponse{Request: response}, nil
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
