package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ClashrAuto/clash-auto/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clash Auto")
	w.Resize(fyne.NewSize(960, 570))

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Status", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Sub", widget.NewLabel("World!")),
	// )

	// tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// tabs.SetTabLocation(container.TabLocationLeading)

	data := []string{"a", "string", "list"}

	w.SetContent(ui.Desktop(data))
	w.ShowAndRun()
}
