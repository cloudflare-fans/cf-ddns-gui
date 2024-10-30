package gui_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/cloudflare-fans/cf-ddns-gui/gui_app/tab"
)

func (_this *App) initMainWindow() {
	a := *_this.mainApp
	w := a.NewWindow("cfDDNS")
	w.Resize(fyne.NewSize(400, 300))
	w.SetFixedSize(true)

	tabs := container.NewAppTabs(
		tab.Home,
		tab.DDNSTasks,
		tab.Logs,
		tab.About,
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	_this.mainWindow = &w
}
