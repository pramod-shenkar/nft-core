package model

import "gorm.io/gorm"

type TokenStatus string

const (
	Minted TokenStatus = "minted"
	Owned  TokenStatus = "owned"
	ToSell TokenStatus = "tosell"
)

type Token struct {
	gorm.Model
	Name         string
	Status       TokenStatus
	CreatorID    string
	OwnerAddress string
	StorageId    string
	FileType     string
	Hash         string
}
