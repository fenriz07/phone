package seeders

import (
	"github.com/fenriz07/phone/students/fenriz/bd/migrations"

	"github.com/jinzhu/gorm"
)

func PhonesLoad(db *gorm.DB) {

	phones := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	//db.Debug().Model(&migrations.Phone{}).Delete(&migrations.Phone{})

	for _, phone := range phones {

		db.Create(&migrations.PhoneNumber{
			Phone: phone,
		})
	}

}
