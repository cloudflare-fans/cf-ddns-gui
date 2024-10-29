package gui_app

import (
	"fyne.io/fyne/v2"
)

type App struct {
	mainApp    *fyne.App
	mainWindow *fyne.Window
	tray       *fyne.Container
}

func (_this *App) Run() {
	(*_this.mainApp).Run()
}
