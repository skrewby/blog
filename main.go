package main

import (
	"log"
	"net/http"

	"github.com/skrewby/blog/generator"
)

func createTestServer(serverPort string) {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	port := ":" + serverPort
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	env := getEnvironmentVariables()

	generatorInit := generator.GeneratorInit{
		RootPath:    "static",
		ContentPath: "content",
	}
	generator := generator.New(generatorInit)
	generator.Generate()

	if env.Debug {
		createTestServer(env.ServerPort)
	}
}
