package gui_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func trayMenu(window *fyne.Window) []*fyne.MenuItem {
	menuItem00 := fyne.NewMenuItem("Show", func() {
		(*window).Show()
	})

	menuItem01 := fyne.NewMenuItem("Manually Execute", func() {
	})

	menuItem02 := fyne.NewMenuItem("Edit Task Config", func() {
	})

	return []*fyne.MenuItem{
		menuItem00,
		menuItem01,
		menuItem02,
	}
}

func (_this *App) InitMainTray() {
	a := *_this.mainGUI
	w := *_this.mainWindow

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}),
			fyne.NewMenuItem("Manually Execute", func() {

			}),
		)
		desk.SetSystemTrayMenu(m)
	}

	w.SetCloseIntercept(func() {
		w.Hide()
	})
}
