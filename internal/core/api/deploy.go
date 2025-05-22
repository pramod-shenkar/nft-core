package api

import (
	"nftledger/common/types"
	"nftledger/internal/adapter/outbound/dlt"

	"github.com/gofiber/fiber/v2/log"
)

type DeployService struct {
	dlt *dlt.Adapter
}

func NewDeployService(
	dlt *dlt.Adapter,
) *DeployService {
	return &DeployService{
		dlt: dlt,
	}
}

func (s *DeployService) Deploy(contract types.Contract) (txnHash string, err error) {

	txnHash, err = s.dlt.DeployContracts(contract)
	if err != nil {
		log.Error(err)
		return
	}
	return txnHash, nil
}
