package main

import (
	"github.com/cloudflare-fans/cf-ddns-gui/gui_app"
)

func main() {
	app := gui_app.App{}
	app.InitMainGUI()
	app.InitMainWindow()
	app.RunTray() // blocking
}
