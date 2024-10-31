package tab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/cloudflare-fans/cf-ddns-gui/resource/emoji/grin"
)

func homeContentBox() *fyne.Container {

	statusEmojiImage := canvas.NewImageFromResource(&fyne.StaticResource{
		StaticName:    "yes",
		StaticContent: grin.Data,
	})
	statusEmojiImage.FillMode = canvas.ImageFillOriginal
	lastScheduledTime := widget.NewLabel("Last Scheduled: 2024-10-31 00:00:00")

	return container.NewVBox(
		statusEmojiImage,
		lastScheduledTime,
	)
}

var Home = container.NewTabItem("Home", homeContentBox())
