package models

import "github.com/jinzhu/gorm"

type Category struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}

func (c *Category) FindAllCategories(db *gorm.DB) (*[]Category, error) {
	var err error
	categories := []Category{}
	err = db.Debug().Model(&Category{}).Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, err
}

