package page

import (
	"fyne.io/fyne/v2"
)

func Welcome(w fyne.Window) fyne.CanvasObject {
	return markdownContent("README.md")
}
