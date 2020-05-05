package main

import (
	db "github.com/fenriz07/phone/students/fenriz/bd"
	"github.com/fenriz07/phone/students/fenriz/bd/migrations"
	"github.com/fenriz07/phone/students/fenriz/bd/seeders"
	"github.com/fenriz07/phone/students/fenriz/normalize"
)

func main() {
	db := db.Start()

	//Iniciamos la migración, con lo cual creamos la tabla.
	migrations.Migrate(db)

	//Cargamos los seeders
	seeders.Insert(db)

	normalize.NormalizePhones(db)

	defer db.Close()
}
