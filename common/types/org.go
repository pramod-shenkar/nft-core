package types

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address
type PrivateKey ecdsa.PrivateKey
type PublicKey ecdsa.PublicKey
