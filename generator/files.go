package generator

import (
	"log"
	"os"
	"path"
)

func createFile(filePath string, fileName string) *os.File {
	fullFilePath := path.Join(filePath, fileName)

	f, err := os.Create(fullFilePath)
	if err != nil {
		log.Fatalf("Failed to create index page: %v", err)
	}
	log.Printf("Created file: %s", fullFilePath)

	return f
}
