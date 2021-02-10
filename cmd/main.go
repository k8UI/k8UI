package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/k8UI/k8UI/internal/ui"
)

func main() {
	layout := ui.NewWindow(app.New())
	layout.App.Run()
}
