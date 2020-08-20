// +build !nogiugui

package main

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/topxeq/qlang"
	specq "github.com/topxeq/qlang/spec"
	"github.com/topxeq/tk"

	// GUI related start

	qlgithub_AllenDang_giu "github.com/topxeq/qlang/lib/github.com/AllenDang/giu"
	qlgithub_AllenDang_giu_imgui "github.com/topxeq/qlang/lib/github.com/AllenDang/giu/imgui"
	// GUI related end
)

func InitGiu() {
	qlang.Import("github_AllenDang_giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("giu", qlgithub_AllenDang_giu.Exports)
	qlang.Import("github_AllenDang_giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)
	qlang.Import("giu_imgui", qlgithub_AllenDang_giu_imgui.Exports)

}

// full version related start

func loadFont() {
	fonts := giu.Context.IO().Fonts()

	rangeVarT := tk.GetVar("FontRange")

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()

	if rangeVarT == nil {
		builder.AddRanges(fonts.GlyphRangesDefault())
	} else {
		rangeStrT := rangeVarT.(string)
		if rangeStrT == "" || tk.StartsWith(rangeStrT, "COMMON") {
			builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
			builder.AddText("辑" + rangeStrT[6:])
		} else if rangeStrT == "FULL" {
			builder.AddRanges(fonts.GlyphRangesChineseFull())
		} else {
			builder.AddText(rangeStrT)
		}
	}

	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/simhei.ttf"

	if tk.Contains(tk.GetOSName(), "rwin") {
		fontPath = "/Library/Fonts/Microsoft/SimHei.ttf"
	}

	fontVarT := tk.GetVar("Font") // "c:/Windows/Fonts/simsun.ttc"

	if fontVarT != nil {
		fontPath = fontVarT.(string)
	}

	fontSizeStrT := "16"

	fontSizeVarT := tk.GetVar("FontSize")

	if fontSizeVarT != nil {
		fontSizeStrT = fontSizeVarT.(string)
	}

	fontSizeT := tk.StrToIntWithDefaultValue(fontSizeStrT, 16)

	// fonts.AddFontFromFileTTF(fontPath, 14)
	fonts.AddFontFromFileTTFV(fontPath, float32(fontSizeT), imgui.DefaultFontConfig, ranges.Data())
}

func loopWindow(windowA *giu.MasterWindow, loopA func()) {
	// wnd := g.NewMasterWindow("Gox Editor", 800, 600, 0, loadFont)

	windowA.Main(loopA)

}

// full version related end

func InitGiuExports() {
	var guiExports = map[string]interface{}{
		"NewMasterWindow":         giu.NewMasterWindow,
		"SingleWindow":            giu.SingleWindow,
		"Window":                  giu.Window,
		"SingleWindowWithMenuBar": giu.SingleWindowWithMenuBar,
		"WindowV":                 giu.WindowV,

		"MasterWindowFlagsNotResizable": giu.MasterWindowFlagsNotResizable,
		"MasterWindowFlagsMaximized":    giu.MasterWindowFlagsMaximized,
		"MasterWindowFlagsFloating":     giu.MasterWindowFlagsFloating,

		// "Layout":          giu.Layout,

		"NewTextureFromRgba": giu.NewTextureFromRgba,

		"Label":                  giu.Label,
		"LabelV":                 giu.LabelV,
		"Line":                   giu.Line,
		"Button":                 giu.Button,
		"InvisibleButton":        giu.InvisibleButton,
		"ImageButton":            giu.ImageButton,
		"InputTextMultiline":     giu.InputTextMultiline,
		"Checkbox":               giu.Checkbox,
		"RadioButton":            giu.RadioButton,
		"Child":                  giu.Child,
		"ComboCustom":            giu.ComboCustom,
		"Combo":                  giu.Combo,
		"ContextMenu":            giu.ContextMenu,
		"Group":                  giu.Group,
		"Image":                  giu.Image,
		"ImageWithFile":          giu.ImageWithFile,
		"ImageWithUrl":           giu.ImageWithUrl,
		"InputText":              giu.InputText,
		"InputTextV":             giu.InputTextV,
		"InputTextFlagsPassword": giu.InputTextFlagsPassword,
		"InputInt":               giu.InputInt,
		"InputFloat":             giu.InputFloat,
		"MainMenuBar":            giu.MainMenuBar,
		"MenuBar":                giu.MenuBar,
		"MenuItem":               giu.MenuItem,
		"PopupModal":             giu.PopupModal,
		"OpenPopup":              giu.OpenPopup,
		"CloseCurrentPopup":      giu.CloseCurrentPopup,
		"ProgressBar":            giu.ProgressBar,
		"Separator":              giu.Separator,
		"SliderInt":              giu.SliderInt,
		"SliderFloat":            giu.SliderFloat,
		"HSplitter":              giu.HSplitter,
		"VSplitter":              giu.VSplitter,
		"TabItem":                giu.TabItem,
		"TabBar":                 giu.TabBar,
		"Row":                    giu.Row,
		"Table":                  giu.Table,
		"FastTable":              giu.FastTable,
		"Tooltip":                giu.Tooltip,
		"TreeNode":               giu.TreeNode,
		"Spacing":                giu.Spacing,
		"Custom":                 giu.Custom,
		"Condition":              giu.Condition,
		"ListBox":                giu.ListBox,
		"DatePicker":             giu.DatePicker,
		"Dummy":                  giu.Dummy,
		// "Widget":             giu.Widget,

		"PrepareMessageBox": giu.PrepareMsgbox,
		"MessageBox":        giu.Msgbox,

		"LoadFont": loadFont,

		"GetConfirm": getConfirmGUI,

		"SimpleInfo":      simpleInfo,
		"SimpleError":     simpleError,
		"SelectFile":      selectFileGUI,
		"SelectSaveFile":  selectFileToSaveGUI,
		"SelectDirectory": selectDirectoryGUI,

		"EditFile":   editFile,
		"LoopWindow": loopWindow,

		"LayoutP": giu.Layout{},

		"Layout": specq.StructOf((*giu.Layout)(nil)),
		"Widget": specq.StructOf((*giu.Widget)(nil)),
	}

	qlang.Import("gui", guiExports)

}
