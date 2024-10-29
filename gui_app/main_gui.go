package gui_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_icon "github.com/cloudflare-fans/cf-ddns-gui/gui_app/icon"
)

func (_this *App) initMainApp() {
	a := app.New()
	icon := fyne.StaticResource{
		StaticName:    "x",
		StaticContent: _icon.Data,
	}
	a.SetIcon(&icon)
	_this.mainApp = &a
}
