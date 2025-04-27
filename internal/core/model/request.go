package model

import (
	"gorm.io/gorm"
)

type RequestStatus string

const (
	Created   RequestStatus = "created"
	Approved  RequestStatus = "approved"
	Rejected  RequestStatus = "rejected"
	Processed RequestStatus = "processed"
	Failed    RequestStatus = "failed"
)

func (r *RequestStatus) FromInt(i int32) {

	switch i {
	case 0:
		*r = Created
	case 1:
		*r = Approved
	case 2:
		*r = Rejected
	case 3:
		*r = Processed
	case 4:
		*r = Failed
	}

}

type Request struct {
	gorm.Model
	Name        string
	Status      RequestStatus
	StorageId   string
	Filetype    string
	FileContent []byte
}
