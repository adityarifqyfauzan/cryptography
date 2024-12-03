package gui

import (
	"testing"

	"fyne.io/fyne/v2/app"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apps := app.New()
			window := apps.NewWindow("test")
			Run(apps, window)
		})
	}
}
