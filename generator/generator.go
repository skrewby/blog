package generator

import (
	"context"
	"log"
	"path"
	"strings"

	"github.com/skrewby/blog/article"
	"github.com/skrewby/blog/layouts"
)

type GeneratorInit struct {
	RootPath    string
	ContentPath string
}

type Generator struct {
	rootPath    string
	contentPath string
	articles    []article.Article
}

func New(init GeneratorInit) *Generator {
	return &Generator{
		rootPath:    init.RootPath,
		contentPath: init.ContentPath,
	}
}

func (g *Generator) Generate() {
	// Create the root folder of our static site (normally)
	createFolder(g.rootPath)

	// Find all folders inside the content folder (these are our articles)
	contentFolders := getSubFolders(g.contentPath)

	// Iterate through all markdown files and generate equivalent html pages
	g.convertContentFiles(contentFolders)

	// Create the index.html page which will be our blog's landing page
	g.createRootPage()

	// Since the program only needs to be run once to generate, we use log.Fatalf to
	// exit the program if anything failed, therefore if it reaches this line, we can
	// assume that everything was executed successfully
	log.Print("Finished generating static site")
}

func (g *Generator) createRootPage() {
	f := createFile(g.rootPath, "index.html")
	component := layouts.Landing(g.articles)
	component.Render(context.Background(), f)
}

func (g *Generator) convertContentFiles(contentFolders []string) {
	for _, folder := range contentFolders {
		contentPath := path.Join(g.contentPath, folder)
		destinationPath := path.Join(g.rootPath, folder)

		createFolder(destinationPath)
		art := article.Article{
			Folder: destinationPath,
		}

		files := getFilesInFolder(contentPath, ".md")
		for _, file := range files {
			destinationFileName, _ := strings.CutSuffix(file, ".md")
			destinationFileName = destinationFileName + ".html"
			meta := g.convert(contentPath, file, destinationPath, destinationFileName)

			page := article.ArticlePage{
				FileName: destinationFileName,
				FullPath: path.Join(folder, destinationFileName),
				Title:    meta.Title,
			}
			art.Pages = append(art.Pages, page)
		}

		g.articles = append(g.articles, art)
	}
}
