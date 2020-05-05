package migrations

import "github.com/jinzhu/gorm"

type PhoneNumber struct {
	gorm.Model
	Phone string
}
