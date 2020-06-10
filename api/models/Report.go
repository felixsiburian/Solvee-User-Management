package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Report struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	UserID      uint32    `json:"user_id"`
	User        User      `json:"user"`
	CategoryID  uint64    `json:"category_id"`
	Category    Category  `json:"category"`
	ProvinceID  uint64    `json:"province_id"`
	Province    Province  `json:"province"`
	RegencyID   uint64    `json:"regency_id"`
	Regency     Regency   `json:"regency"`
	DistrictID  uint32    `json:"district_id"`
	District    Disctrict `json:"district"`
	VillageID   uint64    `json:"village_id"`
	Village     Village   `json:"village"`
	Location    string    `json:"location"`
	CreatedAt   string    `json:"created_at"`
	Description string    `json:"description"`
}

func (r *Report) Prepare () {
	r.ID = 0
	r.Title = html.EscapeString(strings.TrimSpace(r.Title))
	r.User = User{}
	r.Category = Category{}
	r.Province = Province{}
	r.Regency = Regency{}
	r.District = Disctrict{}
	r.Village = Village{}
	r.CreatedAt = html.EscapeString(strings.TrimSpace(r.CreatedAt))
}

func (r *Report) Validate() error {
	if r.Title == "" {
		return errors.New("Required title")
	}
	if r.UserID < 0 {
		return errors.New("Required User")
	}
	if r.CategoryID < 0 {
		return errors.New("Required Category")
	}
	if r.ProvinceID < 0 {
		return errors.New("Required Province")
	}
	if r.RegencyID < 0 {
		return errors.New("Required Regency")
	}
	if r.DistrictID < 0 {
		return errors.New("Required District")
	}
	if r.Location == "" {
		return errors.New("Required Location")
	}
	if r.CreatedAt == "" {
		return errors.New("Required Created Date")
	}
	if r.Description == "" {
		return errors.New("Required Description")
	}
	return nil
}

func (r *Report) SaveReport(db *gorm.DB) (*Report, error) {
	var err error
	err = db.Debug().Model(&Report{}).Create(&r).Error
	if err != nil {
		return &Report{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ? ", r.UserID).Take(&r.User).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Category{}).Where("id = ?", r.CategoryID).Take(&r.Category).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Province{}).Where("id = ? ", r.ProvinceID).Take(&r.Province).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Regency{}).Where("id = ? ", r.RegencyID).Take(&r.Regency).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Disctrict{}).Where("id = ? ", r.DistrictID).Take(&r.District).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Village{}).Where("id = ? ", r.VillageID).Take(&r.Village).Error
		if err != nil {
			return &Report{}, err
		}
	}
	return r , nil
}

func (r *Report) FindAllReport(db *gorm.DB) (*[]Report, error) {
	var err error
	reports := []Report{}
	err = db.Debug().Model(&Report{}).Find(&reports).Error
	if err != nil {
		return &[]Report{}, err
	}
	if len(reports) > 0 {
		for i, _ := range reports {
			err := db.Debug().Model(&User{}).Where("id = ? ", reports[i].UserID).Take(&reports[i].User).Error
			if err != nil {
				return &[]Report{}, err
			}
			err = db.Debug().Model(&Category{}).Where("id = ?", reports[i].CategoryID).Take(&reports[i].Category).Error
			if err != nil {
				return &[]Report{}, err
			}
			err = db.Debug().Model(&Province{}).Where("id = ?", reports[i].ProvinceID).Take(&reports[i].Province).Error
			if err != nil {
				return &[]Report{}, err
			}
			err = db.Debug().Model(&Regency{}).Where("id = ?", reports[i].RegencyID).Take(&reports[i].Regency).Error
			if err != nil {
				return &[]Report{}, err
			}
			err = db.Debug().Model(&Disctrict{}).Where("id = ?", reports[i].DistrictID).Take(&reports[i].District).Error
			if err != nil {
				return &[]Report{}, err
			}
			err = db.Debug().Model(&Village{}).Where("id = ?", reports[i].VillageID).Take(&reports[i].Village).Error
			if err != nil {
				return &[]Report{}, err
			}
		}
	}
	return &reports, nil
}

func (r *Report) FindReportByID (db *gorm.DB, pid uint64) (*Report, error) {
	var err error
	err = db.Debug().Model(&Report{}).Where("id = ?", pid).Take(&r).Error
	if err != nil {
		return &Report{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.UserID).Take(&r.User).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Category{}).Where("id = ?", r.CategoryID).Take(&r.Category).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Province{}).Where("id = ?", r.ProvinceID).Take(&r.Province).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Regency{}).Where("id = ?", r.RegencyID).Take(&r.RegencyID).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Disctrict{}).Where("id = ?", r.DistrictID).Take(&r.District).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Village{}).Where("id = ?", r.VillageID).Take(&r.Village).Error
		if err != nil {
			return &Report{}, err
		}
	}
	return r, nil
}

func (r *Report) UpdateAReport(db *gorm.DB) (*Report, error) {
	var err error

	err = db.Debug().Model(&Report{}).Where("id = ? ", r.ID).Updates(Report{
			Title:       r.Title,
			CategoryID:  r.CategoryID,
			ProvinceID:  r.ProvinceID,
			RegencyID:   r.RegencyID,
			DistrictID:  r.DistrictID,
			VillageID:   r.VillageID,
			Location:    r.Location,
			CreatedAt:   r.CreatedAt,
			Description: r.Description,
		}).Error
	if err != nil {
		return &Report{}, err
	}
	if r.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", r.UserID).Take(&r.User).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Category{}).Where("id = ?", r.CategoryID).Take(&r.Category).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Province{}).Where("id = ?", r.ProvinceID).Take(&r.Province).Error
		if err != nil {
			return &Report{}, err
		}
		//err = db.Debug().Model(&Regency{}).Where("id = ?", r.RegencyID).Take(&r.RegencyID).Error
		//if err != nil {
		//	return &Report{}, err
		//}
		err = db.Debug().Model(&Disctrict{}).Where("id = ?", r.DistrictID).Take(&r.District).Error
		if err != nil {
			return &Report{}, err
		}
		err = db.Debug().Model(&Village{}).Where("id = ?", r.VillageID).Take(&r.Village).Error
		if err != nil {
			return &Report{}, err
		}
	}
	return r, nil
}

func (r *Report) DeleteAReport(db *gorm.DB, pid uint64, uid uint32) (int64, error){
	db = db.Debug().Model(&Report{}).Where("id = ? and user_id = ?", pid, uid).Take(&Report{}).Delete(&Report{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Report Not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}