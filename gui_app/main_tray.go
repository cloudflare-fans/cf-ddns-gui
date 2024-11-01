package gui_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func trayMenu(window *fyne.Window) []*fyne.MenuItem {
	menuItem00 := fyne.NewMenuItem("Show", func() {
		(*window).Show()
	})

	menuSeperator00 := fyne.NewMenuItemSeparator()

	menuItem01 := fyne.NewMenuItem("Manually Execute", func() {
	})

	menuItem02 := fyne.NewMenuItem("Edit Task Config", func() {
	})

	return []*fyne.MenuItem{
		menuItem00,
		menuSeperator00,
		menuItem01,
		menuItem02,
	}
}

func (_this *App) initMainTray() {
	a := *_this.mainApp
	w := *_this.mainWindow

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("trayApp",
			trayMenu(&w)...,
		)
		desk.SetSystemTrayMenu(m)
	}

	w.SetCloseIntercept(func() {
		w.Hide()
	})
}
