package main

import (
	"log"
	"os"
	"strconv"

	"github.com/geoffrey-anto/golang-microservice-apis/server"
	"github.com/joho/godotenv"
)

func getPort() int {
	PORT_s := os.Getenv("PORT")

	if PORT_s == "" {
		log.Fatalf("No PORT env variable")
	}

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Wrong PORT variable")
	}

	return PORT
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	s := server.NewServer("0.0.0.0", getPort())
	s.RunServer()
}
