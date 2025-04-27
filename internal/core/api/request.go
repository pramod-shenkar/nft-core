package api

import (
	"context"
	"nftledger/common/utils"
	"nftledger/internal/adapter/outbound/dao"
	"nftledger/internal/adapter/outbound/storage"
	"nftledger/internal/core/model"
)

type RequestService struct {
	dao     *dao.Daos
	storage *storage.Client
}

func NewRequestService(
	dao *dao.Daos,
	storage *storage.Client,
) *RequestService {
	return &RequestService{
		dao:     dao,
		storage: storage,
	}
}

func (s *RequestService) SaveRequest(request *model.Request) error {

	storageId, err := s.storage.Save(context.Background(), request.FileContent)
	if err != nil {
		return err
	}

	request.Status = model.Created
	request.StorageId = storageId
	err = s.dao.RequestDao.SaveRequest(request)
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

	switch update.Status {
	case model.Approved:

		request, err := s.dao.RequestDao.GetRequest(where)
		if err != nil {
			return err
		}

		payload, err := s.storage.Get(context.Background(), request.StorageId)
		if err != nil {
			return err
		}

		hash, err := utils.ContentHash(payload)
		if err != nil {
			return err
		}

		var token = &model.Token{
			Name:         request.Name,
			Status:       model.Minted,
			CreatorID:    "",
			OwnerAddress: "",
			StorageId:    request.StorageId,
			FileType:     request.Filetype,
			Hash:         hash,
		}

		// todo: instead saving, send it to dlt & then on dltevent save it
		err = s.dao.TokenDao.SaveToken(token)
		if err != nil {
			return err
		}

	}
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
