package generator

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"

	"github.com/skrewby/blog/layouts"
)

type ConvertMeta struct {
	Title   string
	Tags    []string
	Project string
	Page    int
}

func (g Generator) convert(contentFolder string, contentFile string, destinationFolder string, destinationFile string) ConvertMeta {
	f := createFile(destinationFolder, destinationFile)

	contentPath := path.Join(contentFolder, contentFile)
	source, err := os.ReadFile(contentPath)
	if err != nil {
		log.Fatalf("Unable to read content markdown file: %v", err)
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := markdown.Convert([]byte(source), &buf, parser.WithContext(ctx)); err != nil {
		log.Fatalf("Unable to convert markdown file: %v", err)
	}

	metaData := meta.Get(ctx)
	title := fmt.Sprint(metaData["Title"])

	project := fmt.Sprint(metaData["Project"])
	if project == "<nil>" {
		project = ""
	}

	page := 0
	if pageData, ok := metaData["Page"]; ok {
		if pageInt, ok := pageData.(int); ok {
			page = pageInt
		}
	}

	var tags []string
	if tagsData, ok := metaData["Tags"]; ok {
		if tagList, ok := tagsData.([]any); ok {
			for _, tag := range tagList {
				if tagStr, ok := tag.(string); ok {
					tags = append(tags, tagStr)
				}
			}
		}
	}

	convertMeta := ConvertMeta{
		Title:   title,
		Tags:    tags,
		Project: project,
		Page:    page,
	}

	component := layouts.Article(buf.String())
	component.Render(context.Background(), f)

	return convertMeta
}
