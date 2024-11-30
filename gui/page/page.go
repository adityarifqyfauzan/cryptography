package page

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/adityarifqyfauzan/cryptography/docs"
)

func loadMarkdownFromFile(filePath string) string {
	data, err := docs.DocsMarkdown.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func markdownContent(path string) fyne.CanvasObject {
	markdownContent := loadMarkdownFromFile(path)
	markdownView := widget.NewRichTextFromMarkdown(markdownContent)
	markdownView.Segments[0].(*widget.TextSegment).Style.TextStyle.Bold = true
	scrollable := container.NewScroll(markdownView)
	content := container.NewBorder(nil, nil, nil, nil, scrollable)

	return content
}
