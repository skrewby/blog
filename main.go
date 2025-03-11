package main

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/skrewby/blog/layouts"
)

func main() {
	rootPath := "static"
	if err := os.Mkdir(rootPath, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	fileName := path.Join(rootPath, "index.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create index page: %v", err)
	}

	component := layouts.Article("Pedro")
	component.Render(context.Background(), f)
}
