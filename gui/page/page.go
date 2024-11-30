package page

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loadMarkdownFromFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Gagal membaca file Markdown: %v", err)
	}
	return string(content)
}

func markdownContent(path string) fyne.CanvasObject {
	markdownContent := loadMarkdownFromFile(path)
	markdownView := widget.NewRichTextFromMarkdown(markdownContent)
	markdownView.Segments[0].(*widget.TextSegment).Style.TextStyle.Bold = true
	scrollable := container.NewScroll(markdownView)
	content := container.NewBorder(nil, nil, nil, nil, scrollable)

	return content
}
