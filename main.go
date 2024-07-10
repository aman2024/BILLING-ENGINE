package main

import (
	"billing-engine/cron"
	"billing-engine/db"
	"billing-engine/routes"
	"billing-engine/services"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	LoadEnvVariables()

	dbClient, err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	services := services.Services{
		DB: &db.DbClient{Client: dbClient},
	}
	cron.UpdateRepaymentStatusFromLockedToPending(&services)
	cron.UpdateRepaymentStatusFromPendingToSkipped(&services)
	router := routes.Init(&services)
	if err := router.Run(":8000"); err != nil {
		panic(fmt.Sprint("Failed to run routes", err))
	}
}

func LoadEnvVariables() {
	env := os.Getenv("ENV")
	if strings.ToUpper(env) == "DEV" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal("Error loading .env file with error: ", err)
		}
	}
}
