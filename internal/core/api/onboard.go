package api

import (
	"context"
	"nftledger/internal/adapter/outbound/dlt"
	"nftledger/internal/core/model"

	"github.com/gofiber/fiber/v2/log"
)

type OnboardService struct {
	dlt *dlt.Adapter
}

func NewOnboardService(
	dlt *dlt.Adapter,
) *OnboardService {
	return &OnboardService{
		dlt: dlt,
	}
}

func (s *OnboardService) Onboard(request *model.Org) (txnHash string, err error) {
	txnHash, err = s.dlt.OnBoardOrganization(context.Background(), request.Address, request.Role)
	if err != nil {
		log.Error(err)
		return
	}
	return txnHash, nil
}
