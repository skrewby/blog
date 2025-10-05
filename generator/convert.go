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
	Title string
	Tags  []string
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
		panic(err)
	}

	metaData := meta.Get(ctx)
	title := fmt.Sprint(metaData["Title"])

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
		Title: title,
	}

	component := layouts.Article(buf.String())
	component.Render(context.Background(), f)

	return convertMeta
}
