package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment struct {
	Debug      bool
	ServerPort string
}

func getEnvironmentVariables() Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, debugMode := os.LookupEnv("DEBUG_MODE")
	debugStr := os.Getenv("DEBUG_MODE")
	if strings.ToUpper(debugStr) == "FALSE" {
		debugMode = false
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("SERVER_PORT must have a value")
	}

	return Environment{
		Debug:      debugMode,
		ServerPort: serverPort,
	}
}
