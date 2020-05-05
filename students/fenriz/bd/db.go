package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Start() *gorm.DB {
	db, err := gorm.Open("mysql", "fen:S0&7=hKSwyP7ZU#cD{U8]M5:R8U@/gophercises?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return db
}
