// dlt.go

package dlt

import (
	"context"
	"fmt"
	"nftledger/common/config"
	"nftledger/common/constants"
	"nftledger/internal/adapter/inbound/dlt"

	ledgertypes "nftledger/common/types"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type contracts struct {
	requestContract    *RequestContract
	auctionContract    *AuctionContract
	nftContract        *NftContract
	onBoardingContract *OnBoardingContract
}

// Adapter represents the DLT adapter structure
type Adapter struct {
	client            *ethclient.Client
	auth              *bind.TransactOpts
	contracts         contracts
	store             *OrgStore
	deployedContracts map[ledgertypes.Contract]common.Address
}

// NewAdapter creates a new DLT adapter instance
func NewAdapter() (*Adapter, error) {
	// Create authentication
	auth, err := bind.NewKeyedTransactorWithChainID(config.App.Org.PrivateKey, config.App.Dlt.ChainId)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.DialContext(context.Background(), config.App.Dlt.Url)
	if err != nil {
		return nil, err
	}

	// Initialize OnBoarding contract
	onBoardingContract, err := NewOnBoardingContract(config.App.Org.Address, client)
	if err != nil {
		return nil, err
	}

	// Initialize OrgStore contract
	store, err := NewOrgStore(config.App.Org.Address, client)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		client: client,
		auth:   auth,
		contracts: contracts{
			requestContract:    &RequestContract{},
			auctionContract:    &AuctionContract{},
			nftContract:        &NftContract{},
			onBoardingContract: onBoardingContract,
		},
		store: store,
	}, nil
}

// OnBoardOrganization onboards a new organization with a specific role
func (d *Adapter) OnBoardOrganization(ctx context.Context, address common.Address, role uint8) (string, error) {
	// Call the OnBoardOrg method from the OnBoarding contract
	tx, err := d.contracts.onBoardingContract.OnBoardOrg(d.auth, address, role)
	if err != nil {
		return "", err
	}

	// Wait for the transaction to be mined
	_, err = bind.WaitMined(ctx, d.client, tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

// GetOrganizationRole retrieves the role of an organization
func (d *Adapter) GetOrganizationRole(ctx context.Context, address common.Address) (uint8, error) {
	// Call the GetOrg method from the OrgStore contract
	role, err := d.store.GetOrg(&bind.CallOpts{
		Context: ctx,
	}, address)
	if err != nil {
		return 0, err
	}

	return role, nil
}

// SaveOrganization saves an organization's information
func (d *Adapter) SaveOrganization(ctx context.Context, address common.Address, role uint8) error {
	// Call the SaveOrg method from the OrgStore contract
	tx, err := d.store.SaveOrg(d.auth, address, role)
	if err != nil {
		return err
	}

	// Wait for the transaction to be mined
	_, err = bind.WaitMined(ctx, d.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// WatchOnBoardEvents watches for OnBoard events
func (d *Adapter) WatchOnBoardEvents(ctx context.Context, eventChan chan<- *OnBoardingContractOnBoardEvent) error {
	// Create watch options
	watchOpts := &bind.WatchOpts{
		Context: ctx,
	}

	// Start watching for OnBoard events
	sub, err := d.contracts.onBoardingContract.WatchOnBoardEvent(watchOpts, eventChan)
	if err != nil {
		return err
	}

	// Handle subscription errors
	go func() {
		for {
			select {
			case err := <-sub.Err():
				if err != nil {
					// Handle error appropriately
					return
				}
			case <-ctx.Done():
				sub.Unsubscribe()
				return
			}
		}
	}()

	return nil
}

func (d *Adapter) DeployContracts(contract ledgertypes.Contract) (string, error) {

	var contractAddress common.Address
	var txn *types.Transaction
	var err error

	switch contract {
	case constants.RequestContract:
		contractAddress, txn, _, err = dlt.DeployRequestContract(d.auth, d.client)
		if err != nil {
			return "", err
		}
		d.deployedContracts[constants.RequestContract] = contractAddress
	case constants.AuctionContract:
		contractAddress, txn, _, err = dlt.DeployAuctionContract(d.auth, d.client)
		if err != nil {
			return "", err
		}
		d.deployedContracts[constants.AuctionContract] = contractAddress
	case constants.NftContract:
		contractAddress, txn, _, err = dlt.DeployNftContract(d.auth, d.client)
		if err != nil {
			return "", err
		}
		d.deployedContracts[constants.NftContract] = contractAddress
	case constants.OnboardingContract:
		contractAddress, txn, _, err = dlt.DeployOnBoardingContract(d.auth, d.client)
		if err != nil {
			return "", err
		}
		d.deployedContracts[constants.OnboardingContract] = contractAddress
	default:
		return "", fmt.Errorf("contract %s not found", contract)
	}

	return txn.Hash().Hex(), nil
}
