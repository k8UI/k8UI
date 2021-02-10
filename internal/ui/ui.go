package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/vrischmann/envconfig"

	"github.com/k8UI/k8UI/internal/config"
	"github.com/k8UI/k8UI/internal/ui/subsections"
)

const AppName = "k8UI"

type WindowLayout struct {
	App    fyne.App
	config config.Config
}

func NewWindow(a fyne.App) *WindowLayout {
	cfg := config.Config{}
	if err := envconfig.Init(&cfg); err != nil {
		panic(err)
	}

	w := a.NewWindow(AppName)
	content := container.NewMax()
	content.Add(widget.NewButton("Test", func() {}))
	split := container.NewHSplit(makeNav(), content)
	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 460))
	w.Show()

	return &WindowLayout{
		App:    a,
		config: cfg,
	}
}

func makeNav() fyne.CanvasObject {
	tree := &widget.Tree{
		ChildUIDs: func(uid string) (c []string) {
			c = subsections.SubsIndex[uid]
			return
		},
		IsBranch: func(uid string) (b bool) {
			_, b = subsections.SubsIndex[uid]
			return
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		UpdateNode: func(uid string, branch bool, node fyne.CanvasObject) {
			node.(*widget.Label).SetText(uid)
		},
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}
