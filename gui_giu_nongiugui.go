// +build !giugui

package main

import (
	"os"

	"github.com/topxeq/qlang"
	"github.com/topxeq/tk"

	// GUI related start
	qlfyne_fyne "github.com/topxeq/qlang/lib/fyne.io/fyne"
	qlfyne_fyne_app "github.com/topxeq/qlang/lib/fyne.io/fyne/app"
	qlfyne_fyne_canvas "github.com/topxeq/qlang/lib/fyne.io/fyne/canvas"
	qlfyne_fyne_dialog "github.com/topxeq/qlang/lib/fyne.io/fyne/dialog"
	qlfyne_fyne_layout "github.com/topxeq/qlang/lib/fyne.io/fyne/layout"
	qlfyne_fyne_widget "github.com/topxeq/qlang/lib/fyne.io/fyne/widget"
	// GUI related end
)

func InitGiu() {
	// GUI related start
	qlang.Import("fyne", qlfyne_fyne.Exports)
	qlang.Import("fyne_app", qlfyne_fyne_app.Exports)
	qlang.Import("fyne_widget", qlfyne_fyne_widget.Exports)
	qlang.Import("fyne_canvas", qlfyne_fyne_canvas.Exports)
	qlang.Import("fyne_dialog", qlfyne_fyne_dialog.Exports)
	qlang.Import("fyne_layout", qlfyne_fyne_layout.Exports)

	fontPathT := tk.Trim(os.Getenv("FYNE_FONT"))

	if fontPathT == "" {
		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			os.Setenv("FYNE_FONT", `/usr/share/fonts/SimHei.ttf`)
		} else if tk.Contains(osT, "arwin") {
			os.Setenv("FYNE_FONT", `/Library/Fonts/Microsoft/SimHei.ttf`)
		} else {
			os.Setenv("FYNE_FONT", `c:/Windows/Fonts/simsun.ttc`)
		}
	}

	// GUI related end

}

// full version related start

func loadFont() {
}

// full version related end

func InitGiuExports() {

}
