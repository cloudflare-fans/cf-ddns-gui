package gui_app

import "fyne.io/fyne/v2/app"

func (_this *App) InitMainGUI() {
	a := app.New()
	_this.mainApp = &a
}
