package tab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_icon "github.com/cloudflare-fans/cf-ddns-gui/gui_app/icon"
	"image/color"
	_url "net/url"
)

func aboutContentBox() *fyne.Container {
	title := canvas.NewText("cfDDNS", color.Black)
	title.TextSize = 25
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true

	description := widget.NewLabelWithStyle(
		"An infamous method on scheduling DDNS tasks\n"+
			"to your domains on Cloudflare dashboard.",
		fyne.TextAlignCenter, fyne.TextStyle{Bold: true},
	)

	url, _ := _url.Parse("https://github.com/cloudflare-fans/cf-ddns")
	repoLink := widget.NewHyperlinkWithStyle("Github", url, fyne.TextAlignCenter, fyne.TextStyle{})

	url, _ = _url.Parse("https://github.com/cloudflare-fans/cf-ddns/blob/main/LICENSE")
	licenseLink := widget.NewHyperlinkWithStyle("Apache License 2.0", url, fyne.TextAlignCenter, fyne.TextStyle{})

	url, _ = _url.Parse("https://github.com/cloudflare-fans")
	homepageLink := widget.NewHyperlinkWithStyle("© Cloudflare Fans", url, fyne.TextAlignCenter, fyne.TextStyle{})

	icon := fyne.StaticResource{
		StaticName:    "cfDDNS logo",
		StaticContent: _icon.Data,
	}
	logoIcon := canvas.NewImageFromResource(&icon)
	logoIcon.FillMode = canvas.ImageFillOriginal

	return container.New(
		layout.NewCustomPaddedLayout(16, 16, 16, 16),
		container.NewVBox(
			container.New(layout.NewCustomPaddedHBoxLayout(16), logoIcon, container.NewVBox(
				title,
				description,
			)),
			container.New(layout.NewCenterLayout(), container.NewHBox(repoLink, licenseLink, homepageLink)),
		),
	)
}

var About = container.NewTabItem("About", aboutContentBox())
