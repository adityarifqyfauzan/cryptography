package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func Run(a fyne.App, w fyne.Window) {
	a.SetIcon(data.FyneLogo)
	logLifecycle(a)
	topWindow = w
	w.SetMaster()

	// Cache untuk menyimpan tampilan menu
	viewCache := make(map[string]fyne.CanvasObject)

	content := container.NewStack()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord

	setTutorial := func(t Menu) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		if t.Title == "Welcome" {
			title.Hide()
			intro.Hide()
		} else {
			title.Show()
			intro.Show()
		}

		// Cek apakah tampilan menu sudah ada di cache
		view, exists := viewCache[t.Title]
		if !exists {
			// Jika belum ada di cache, buat tampilan baru
			view = t.View(w)
			viewCache[t.Title] = view
		}

		// Ganti isi konten tanpa memanggil `content.Refresh`
		content.Objects = []fyne.CanvasObject{view}
	}

	menu := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(makeNav(setTutorial, false))
	} else {
		split := container.NewHSplit(makeNav(setTutorial, true), menu)
		split.Offset = 0.2
		w.SetContent(split)
	}
	w.Resize(fyne.NewSize(640, 460))
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeNav(setMenu func(menu Menu), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()
	var tree *widget.Tree
	tree = &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return MenuIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := MenuIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := Menus[uid]
			if !ok {
				fyne.LogError("Missing menu panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Menus[uid]; ok {
				// Toggle dropdown expand state
				if tree.IsBranch(uid) {
					if tree.IsBranchOpen(uid) {
						tree.CloseBranch(uid) // Tutup jika sudah terbuka
					} else {
						tree.OpenBranch(uid) // Buka jika tertutup
					}
				}

				// Set menu konten
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setMenu(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
