package api

import (
	"github.com/joho/godotenv"
	"go-blog/api/controllers"
	"log"
	"os"
)

var server = controllers.Server{}

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file not found")
	}
}

func Run()  {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(":8081")
}
