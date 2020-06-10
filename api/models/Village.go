package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Village struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	District   Disctrict `json:"district"`
	DistrictID uint32    `json:"district_id"`
	Name       string    `json:"name"`
}

func (v *Village) FindAllVillage(db *gorm.DB) (*[]Village, error) {
	var err error
	villages := []Village{}
	err = db.Debug().Model(&Village{}).Find(&villages).Error
	if err != nil {
		return &[]Village{}, err
	}
	if len(villages) > 0 {
		for i, _ := range villages {
			err := db.Debug().Model(&Village{}).Where("id = ? ", villages[i].DistrictID).Take(&villages[i].District).Error
			if err != nil {
				return &[]Village{}, err
			}
		}
	}
	return &villages, nil
}

func (v *Village) FindVillageByID (db *gorm.DB, pid uint64) (*Village, error){
	var err error
	err = db.Debug().Model(&Village{}).Where("id = ?", pid).Take(&v).Error
	if err != nil {
		return &Village{}, err
	}
	if v.ID != 0 {
		err = db.Debug().Model(&Disctrict{}).Where("id = ?", v.DistrictID).Take(&v.District).Error
		if err != nil {
			return &Village{}, err
		}
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Village{}, errors.New("Village Not Found")
	}
	return v, nil
}
