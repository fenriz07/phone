package migrations

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {

	db.DropTableIfExists(&PhoneNumber{})

	db.AutoMigrate(&PhoneNumber{})
}
