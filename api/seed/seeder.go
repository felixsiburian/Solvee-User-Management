package seed

import (
	"Solvee-User-Management/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

//memasukkan data dummy ke database agar database tidak kosong, sehingga bisa menggunakan API get User dahulu

var users = []models.User{
	models.User{
		Name:        "User",
		Email:       "user@gmail.com",
		PhoneNumber: "08123222312",
		Address:     "Jalan jalan",
		Password:    "password",
	},
	models.User{
		Name:        "User 2",
		Email:       "user2@gmail.com",
		PhoneNumber: "089978662232",
		Address:     "Jalan kaki",
		Password:    "password",
	},
}

func Load(db *gorm.DB){
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table : %v", err)
		}
	}

}