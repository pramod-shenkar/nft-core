package api

import (
	"nftledger/internal/adapter/outbound/dao"
	"nftledger/internal/core/model"
)

type RequestService struct {
	dao *dao.Daos
}

func NewRequestService(
	dao *dao.Daos,
) *RequestService {
	return &RequestService{
		dao: dao,
	}
}

func (s *RequestService) SaveRequest(request *model.Request) error {
	err := s.dao.RequestDao.SaveRequest(request)
	if err != nil {
		return err
	}
	return nil
}

func (s *RequestService) GetRequest(request *model.Request) (*model.Request, error) {
	request, err := s.dao.RequestDao.GetRequest(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *RequestService) GetRequests(request *model.Request) (*[]model.Request, error) {
	var requests *[]model.Request
	requests, err := s.dao.RequestDao.GetRequests(request)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (s *RequestService) UpdateRequest(where, update *model.Request) error {
	err := s.dao.RequestDao.UpdateRequest(where, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *RequestService) DeleteRequest(request *model.Request) error {
	err := s.dao.RequestDao.DeleteRequest(request)
	if err != nil {
		return err
	}
	return nil
}
