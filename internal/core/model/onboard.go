package model

import "github.com/ethereum/go-ethereum/common"

type Org struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Address common.Address `json:"address"`
	Role    uint8          `json:"role"`
}
