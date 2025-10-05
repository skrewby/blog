package generator

import (
	"log"
	"os"
	"strings"
)

func createFolder(name string) {
	if err := os.Mkdir(name, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create directory (%s): %v", name, err)
	}

	log.Printf("Created folder: %s", name)
}

func getSubFolders(parent string) []string {
	entries, err := os.ReadDir(parent)
	if err != nil {
		log.Fatalf("Failed to read folder (%s): %v", parent, err)
	}

	var folders []string
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	return folders
}

func getFilesInFolder(folder string, ext string) []string {
	var filesFound []string

	log.Printf("Getting files with extension %s in folder %s", ext, folder)
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("Failed to read directory (%s): %v", folder, err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ext) {
			continue
		}

		log.Printf("Found file: %s", file.Name())
		filesFound = append(filesFound, file.Name())
	}

	return filesFound
}
