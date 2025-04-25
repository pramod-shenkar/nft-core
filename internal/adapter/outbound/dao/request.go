package dao

import (
	"nftledger/internal/adapter/outbound/dao/db"
	"nftledger/internal/core/model"
)

type RequestDao struct {
	*db.Conn
}

func NewRequestDao(db *db.Conn) *RequestDao {
	return &RequestDao{
		db,
	}
}

func (d *RequestDao) SaveRequest(request *model.Request) error {
	err := d.Conn.DB.Model(&model.Request{}).Create(request).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *RequestDao) GetRequest(request *model.Request) (*model.Request, error) {
	err := d.Conn.DB.Model(&model.Request{}).Where(request).First(&request).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (d *RequestDao) GetRequests(request *model.Request) (*[]model.Request, error) {
	var requests []model.Request
	err := d.Conn.DB.Model(&model.Request{}).Where(request).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return &requests, nil
}

func (d *RequestDao) UpdateRequest(where, update *model.Request) error {
	err := d.Conn.DB.Model(&model.Request{}).Where(where).Updates(&update).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *RequestDao) DeleteRequest(request *model.Request) error {
	err := d.Conn.DB.Model(&model.Request{}).Delete(request).Error
	if err != nil {
		return err
	}
	return nil
}
