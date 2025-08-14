package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/skrewby/blog/layouts"
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

func main() {
	env := getEnvironmentVariables()

	rootPath := "static"
	if err := os.Mkdir(rootPath, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	fileName := path.Join(rootPath, "index.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create index page: %v", err)
	}

	component := layouts.Article("Pedro")
	component.Render(context.Background(), f)

	if env.Debug {
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/", fs)
		port := ":" + env.ServerPort
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
