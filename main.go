package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tatthien/go-cms-api/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	ip := os.Getenv("APP_IP")
	port := os.Getenv("APP_PORT")

	server := server.New(ip, port)
	server.Run()
}
