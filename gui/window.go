package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Desktop(data []string) *fyne.Container {

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Loading...")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		},
	)

	out := container.NewVBox(container.NewAppTabs(
		container.NewTabItem("Status", container.NewVBox(
			list,
		)),
		container.NewTabItem("Sub", widget.NewLabel("World!")),
	))
	// in := widget.NewEntry()

	// in.OnChanged = func(content string) {
	// 	out.SetText("Hello " + content + "!")
	// }
	return out
}
