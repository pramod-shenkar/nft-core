package dao

import (
	"nftledger/internal/adapter/outbound/dao/db"
	"nftledger/internal/core/model"
)

type TokenDao struct {
	*db.Conn
}

func NewTokenDao(db *db.Conn) *TokenDao {
	return &TokenDao{
		db,
	}
}

func (d *TokenDao) SaveToken(token *model.Token) error {
	err := d.Conn.DB.Model(&model.Token{}).Create(token).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *TokenDao) GetToken(token *model.Token) (*model.Token, error) {
	err := d.Conn.DB.Model(&model.Token{}).Where(token).First(&token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (d *TokenDao) GetTokens(token *model.Token) (*[]model.Token, error) {
	var tokens []model.Token
	err := d.Conn.DB.Model(&model.Token{}).Where(token).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return &tokens, nil
}

func (d *TokenDao) UpdateToken(where, update *model.Token) error {
	err := d.Conn.DB.Model(&model.Token{}).Where(where).Save(&update).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *TokenDao) DeleteToken(token *model.Token) error {
	err := d.Conn.DB.Model(&model.Token{}).Where(&model.Token{Name: token.Name}).Delete(token).Error
	if err != nil {
		return err
	}
	return nil
}
