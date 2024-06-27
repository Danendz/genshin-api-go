package main

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	creds := db.NewDBCreds()
	client, err := db.NewMongoDB(creds)

	if err != nil {
		log.Fatal(err)
	}

	api.NewFiberApp(client, creds)
}
