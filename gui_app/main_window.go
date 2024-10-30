package gui_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/cloudflare-fans/cf-ddns-gui/gui_app/tab"
	"image/color"
)

func (_this *App) initMainWindow() {
	a := *_this.mainApp
	w := a.NewWindow("cfDDNS")
	w.Resize(fyne.NewSize(400, 400))
	w.SetFixedSize(true)

	tabs := container.NewAppTabs(
		tab.Home,
		tab.DDNS,
		tab.Logs,
		tab.About,
	)

	tabs.SetTabLocation(container.TabLocationTop)

	version := canvas.NewText("Version: 0.0.1", color.Black)
	version.Alignment = fyne.TextAlignTrailing

	w.SetContent(
		container.NewBorder(
			nil,
			container.NewVBox(
				canvas.NewLine(color.RGBA{R: 128, G: 128, B: 128, A: 255}),
				version,
			),
			nil,
			nil,
			tabs,
		),
	)
	_this.mainWindow = &w
	w.Show()
}
