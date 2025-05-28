package config

import (
	"crypto/ecdsa"
	"math/big"
	"nftledger/common/utils"
	"strings"

	"github.com/gofiber/fiber/v2/log"

	"github.com/caarlos0/env/v11"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

type orgType string

const (
	Deployer      = "deployer"
	Initiator     = "initiator"
	Transactioner = "transactioner"
	Auditor       = "auditor"
)

type Env string

const (
	Dev  Env = "dev"
	Test Env = "test"
	Prod Env = "prod"
)

type endPointType string

const (
	Rest endPointType = "rest"
	Grpc endPointType = "grpc"
)

type config struct {
	Service struct {
		Env          Env          `env:"env" envDefault:"dev"`
		Port         string       `env:"port" envDefault:"3000"`
		EndPointType endPointType `env:"port" envDefault:"grpc"`
	}
	Org struct {
		Name       string  `env:"OrgName" envDefault:"Deployer01"`
		Type       orgType `env:"OrgType" envDefault:"deployer"`
		privateKey string  `env:"PrivateKey"`
		publicKey  string  `env:"PublicKey"`
		Id         int     `env:"Id" envDefault:"1"`
		address    string  `env:"Address"`
		Address    common.Address
		PrivateKey *ecdsa.PrivateKey
		PublicKey  *ecdsa.PublicKey
	} `env:"Org"`
	Dlt struct {
		Url     string `env:"DltUrl"  envDefault:"http://localhost:9474"`
		chainId int    `env:"ChainId" envDefault:"9474"`
		ChainId *big.Int
	} `env:"Dlt"`
}

func New() (*config, error) {

	godotenv.Load(".env")

	var config config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}

	if config.Service.Env == Dev {
		err = utils.InitOrgsForDevEnv()
		if err != nil {
			return nil, err
		}

		_, config.Org.privateKey, config.Org.publicKey, err = utils.GetOrgKeys(config.Org.Name)
		if err != nil {
			return nil, err
		}
	}

	raw := strings.TrimPrefix(config.Org.privateKey, "0x")
	config.Org.PrivateKey, err = crypto.HexToECDSA(raw)
	if err != nil {
		return nil, err
	}

	config.Org.PublicKey = &config.Org.PrivateKey.PublicKey
	config.Org.Address = crypto.PubkeyToAddress(config.Org.PrivateKey.PublicKey)

	config.Dlt.ChainId = big.NewInt(9474)
	return &config, nil
}

var App *config

func init() {
	config, err := New()
	if err != nil {
		log.Fatal(err)
	}
	App = config
}
