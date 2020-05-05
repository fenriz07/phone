package normalize

import (
	"strings"

	"github.com/fenriz07/phone/students/fenriz/bd/migrations"
	"github.com/jinzhu/gorm"
)

func NormalizePhones(db *gorm.DB) {

	numbers := []migrations.PhoneNumber{}

	//Obtenemos todos los registros de la tabla
	db.Find(&numbers)

	for _, v := range numbers {
		phone := v.Phone
		phone = strings.ReplaceAll(phone, " ", "")
		phone = strings.ReplaceAll(phone, "(", "")
		phone = strings.ReplaceAll(phone, ")", "")
		phone = strings.ReplaceAll(phone, "-", "")

		v.Phone = phone
		db.Save(&v)
	}

	numbers = []migrations.PhoneNumber{}

	db.Find(&numbers)

	idsMark := make(map[uint]bool)

	for _, v := range numbers {

		id := v.ID

		if _, ok := idsMark[id]; ok {
			continue
		}

		phoneField := v.Phone

		phones := []migrations.PhoneNumber{}

		db.Where("phone = ?", phoneField).Find(&phones)

		for _, p := range phones {

			if id != p.ID {
				idsMark[p.ID] = true
				db.Delete(&p)
			}
		}

	}

}
