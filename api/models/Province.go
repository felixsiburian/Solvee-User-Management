package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Province struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}

//Get Smua Provinsi
func (p *Province) FindAllProvince(db *gorm.DB) (*[]Province, error) {
	var err error
	provinces := []Province{}
	err = db.Debug().Model(&Province{}).Find(&provinces).Error
	if err != nil {
		return &[]Province{}, err
	}

	return &provinces, err
}

//Get Provinsi berdasarkan ID
func (p *Province) FindProvinceByID(db *gorm.DB, pid uint64) (*Province, error) {
	var err error
	err = db.Debug().Model(Province{}).Where("id = ? ", pid).Take(&p).Error
	if err != nil {
		return &Province{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Province{}, errors.New("Province Not Found")
	}
	return p, err
}
