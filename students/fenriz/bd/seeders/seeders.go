package seeders

import (
	"github.com/jinzhu/gorm"
)

func Insert(db *gorm.DB) {
	PhonesLoad(db)
}
