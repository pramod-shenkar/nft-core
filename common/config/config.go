package config

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/caarlos0/env/v11"
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

type config struct {
	Org struct {
		Name       string  `env:"OrgName"`
		Type       orgType `env:"OrgType"`
		privateKey string  `env:"PrivateKey"`
		publicKey  string  `env:"PublicKey"`
		Id         int     `env:"Id"`
		Address    string  `env:"Address"`
		PrivateKey *ecdsa.PrivateKey
		PublicKey  *ecdsa.PublicKey
	} `env:"Org"`
	Dlt struct {
		Url     string `env:"DltUrl"`
		chainId int    `env:"ChainId"`
		ChainId *big.Int
	} `env:"Dlt"`
}

func New() (*config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var config config
	err = env.Parse(&config)
	if err != nil {
		return nil, err
	}

	config.Org.PrivateKey, err = crypto.HexToECDSA(config.Org.privateKey)
	if err != nil {
		return nil, err
	}
	config.Org.PublicKey = &config.Org.PrivateKey.PublicKey
	config.Org.Address = crypto.PubkeyToAddress(config.Org.PrivateKey.PublicKey).Hex()

	config.Dlt.ChainId = big.NewInt(int64(config.Dlt.chainId))
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
