package database

import (
	"go-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}

	DB.AutoMigrate(&models.Personalidade{})
}
