package tab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_url "net/url"
)

func aboutContentBox() *fyne.Container {
	url, _ := _url.Parse("https://github.com/cloudflare-fans/cf-ddns")
	var box = container.NewVBox(
		widget.NewLabelWithStyle("cfDDNS", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:   true,
			Italic: true,
		}),
		widget.NewHyperlinkWithStyle("Home Page", url, fyne.TextAlignCenter, fyne.TextStyle{}),
	)
	return box
}

var About = container.NewTabItem("About", aboutContentBox())
