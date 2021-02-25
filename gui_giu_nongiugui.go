// +build !giugui

package main

import (
	"os"

	"github.com/topxeq/qlang"
	"github.com/topxeq/tk"

	// GUI related start
	qlfyne_fyne "github.com/topxeq/qlang/lib/fyne.io/fyne/v2"
	qlfyne_fyne_app "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/app"
	qlfyne_fyne_canvas "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/canvas"
	qlfyne_fyne_container "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/container"
	qlfyne_fyne_dialog "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/dialog"
	qlfyne_fyne_layout "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/layout"
	qlfyne_fyne_widget "github.com/topxeq/qlang/lib/fyne.io/fyne/v2/widget"
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
	qlang.Import("fyne_container", qlfyne_fyne_container.Exports)

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
