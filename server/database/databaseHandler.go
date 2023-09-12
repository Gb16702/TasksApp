package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"todoapp/database/models"
)

var DB *gorm.DB

func ConnectToDatabase(connectionString string) {
	var err error;

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database");
	}

	DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Connecté à la base de données")


	err = DB.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		panic("Could not migrate the database");
	}

	log.Println("Schémas migrés avec succès")
}
