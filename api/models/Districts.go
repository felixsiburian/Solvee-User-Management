package models

import (
	"github.com/jinzhu/gorm"
)

type Disctrict struct {
	ID        uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Regency   Regency `json:"regency"`
	RegencyID uint64  `json:"regency_id"`
	Name      string  `json:"name"`
}

func (d *Disctrict) FindAllDistrict(db *gorm.DB) (*[]Disctrict, error) {
	var err error
	districts := []Disctrict{}
	err = db.Debug().Model(&Disctrict{}).Find(&districts).Error
	if err != nil {
		//return &[]Disctrict{}, err
		return nil, nil
	}
	if len(districts) > 0 {
		for i, _ := range districts {
			err := db.Debug().Model(&Regency{}).Where("id = ?", districts[i].RegencyID).Take(&districts[i].Regency).Error
			if err != nil {
				//return &[]Disctrict{}, err
				return nil, nil
			}
		}
	}

	return &districts, nil
}

func (d *Disctrict) FindDistrictByID(db *gorm.DB, pid uint64) (*Disctrict, error) {
	var err error
	err = db.Debug().Model(&Disctrict{}).Where("id = ?", pid).Take(&d).Error
	if err != nil {
		//return &[]Disctrict{}, err
		return nil, nil
	}

	if d.ID != 0 {
		err = db.Debug().Model(&Regency{}).Where("id = ?", d.RegencyID).Take(&d.RegencyID).Error
		if err != nil {
			//return &[]Disctrict{}, err
			return nil, nil
		}
	}
	return d , nil
}