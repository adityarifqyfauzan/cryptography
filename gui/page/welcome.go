package page

import (
	"fyne.io/fyne/v2"
)

func Welcome(w fyne.Window) fyne.CanvasObject {
	content := markdownContent("README.md")
	return content
}
