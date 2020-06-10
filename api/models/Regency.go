package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Regency struct {
	ID         uint64   `gorm:"primary_key;auto_increment" json:"id"`
	Province   Province `json:"province"`
	ProvinceID uint64   `json:"province_id"`
	Name       string   `json:"name"`
}

func (r *Regency) FindALLRegency(db *gorm.DB) (*[]Regency, error) {
	var err error
	regencies := []Regency{}
	err = db.Debug().Model(&Regency{}).Find(&regencies).Error
	if err != nil {
		return &[]Regency{}, err
	}
	if len(regencies) > 0 {
		for i, _ := range regencies {
			err := db.Debug().Model(&Province{}).Where("id = ?", regencies[i].ProvinceID).Take(&regencies[i].Province).Error
			if err != nil {
				return &[]Regency{}, err
			}
		}
	}

	return &regencies, nil
}

func (r *Regency) FindRegencyByID(db *gorm.DB, pid uint64) (*Regency, error) {
	var err error
	err = db.Debug().Model(&Regency{}).Where("id = ?", pid).Take(&r).Error
	if err != nil {
		return &Regency{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&Province{}).Where("id = ?", r.ProvinceID).Take(&r.Province).Error
		if err != nil {
			return &Regency{}, err
		}
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Regency{}, errors.New("Regency Not Found")
	}
	return r, nil
}
