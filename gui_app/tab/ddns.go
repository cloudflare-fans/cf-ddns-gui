package tab

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{
	{"top left", "top right"},
	{"bottom left", "bottom right"},
}

func ddnsContextBox() *fyne.Container {
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	return container.NewVBox(list)
}

var DDNS = container.NewTabItem("DDNS", ddnsContextBox())
