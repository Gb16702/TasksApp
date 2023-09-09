package main

import (
	"os"
	server "todoapp/core"
	"todoapp/database"

	"github.com/joho/godotenv"
)

type environment struct {
	PORT 		string
	DB_URL		string
}

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Could not load the environment variables")
	}

	environmentVariables := environment{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}

	server.HandleServerStart(
		environmentVariables.PORT,
		environmentVariables.DB_URL,
	);

	database.ConnectToDatabase(environmentVariables.DB_URL)
}
