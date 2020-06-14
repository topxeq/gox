// +build windows

package main

import (
	"runtime"

	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/api"
	"github.com/topxeq/govcl/vcl/rtl"
	"github.com/topxeq/govcl/vcl/types"
	"github.com/topxeq/qlang"
	"github.com/topxeq/tk"

	execq "github.com/topxeq/qlang/exec"

	qlgithub_topxeq_govcl_vcl "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl"
	qlgithub_topxeq_govcl_vcl_api "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/api"
	qlgithub_topxeq_govcl_vcl_rtl "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/rtl"
	qlgithub_topxeq_govcl_vcl_types "github.com/topxeq/qlang/lib/github.com/topxeq/govcl/vcl/types"
)

func InitLCLFirst() {
	var lclExports = map[string]interface{}{
		// "NewTNotifyEvent":                      NewTNotifyEvent,
		// "NewTKeyEvent":                         NewTKeyEvent,
		"NewTNotifyEvent":                      NewTNotifyEvent,
		"NewTKeyEvent":                         NewTKeyEvent,
		"NewTKeyPressEvent":                    NewTKeyPressEvent,
		"NewTMouseEvent":                       NewTMouseEvent,
		"NewTMouseMoveEvent":                   NewTMouseMoveEvent,
		"NewTExceptionEvent":                   NewTExceptionEvent,
		"NewTCloseEvent":                       NewTCloseEvent,
		"NewTCloseQueryEvent":                  NewTCloseQueryEvent,
		"NewTContextPopupEvent":                NewTContextPopupEvent,
		"NewTDragDropEvent":                    NewTDragDropEvent,
		"NewTDragOverEvent":                    NewTDragOverEvent,
		"NewTStartDragEvent":                   NewTStartDragEvent,
		"NewTEndDragEvent":                     NewTEndDragEvent,
		"NewTAlignPositionEvent":               NewTAlignPositionEvent,
		"NewTDockDropEvent":                    NewTDockDropEvent,
		"NewTDockOverEvent":                    NewTDockOverEvent,
		"NewTStartDockEvent":                   NewTStartDockEvent,
		"NewTUnDockEvent":                      NewTUnDockEvent,
		"NewTGetSiteInfoEvent":                 NewTGetSiteInfoEvent,
		"NewTMouseWheelEvent":                  NewTMouseWheelEvent,
		"NewTMouseWheelUpDownEvent":            NewTMouseWheelUpDownEvent,
		"NewTMessageEvent":                     NewTMessageEvent,
		"NewTHelpEvent":                        NewTHelpEvent,
		"NewTWebTitleChangeEvent":              NewTWebTitleChangeEvent,
		"NewTWebJSExternalEvent":               NewTWebJSExternalEvent,
		"NewTMeasureItemEvent":                 NewTMeasureItemEvent,
		"NewTMovedEvent":                       NewTMovedEvent,
		"NewTDrawCellEvent":                    NewTDrawCellEvent,
		"NewTSelectCellEvent":                  NewTSelectCellEvent,
		"NewTGetEditEvent":                     NewTGetEditEvent,
		"NewTSetEditEvent":                     NewTSetEditEvent,
		"NewTDropFilesEvent":                   NewTDropFilesEvent,
		"NewTConstrainedResizeEvent":           NewTConstrainedResizeEvent,
		"NewTWndProcEvent":                     NewTWndProcEvent,
		"NewTSectionNotifyEvent":               NewTSectionNotifyEvent,
		"NewTSectionTrackEvent":                NewTSectionTrackEvent,
		"NewTSectionDragEvent":                 NewTSectionDragEvent,
		"NewTSysLinkEvent":                     NewTSysLinkEvent,
		"NewTDrawItemEvent":                    NewTDrawItemEvent,
		"NewTLVSelectItemEvent":                NewTLVSelectItemEvent,
		"NewTLVCheckedItemEvent":               NewTLVCheckedItemEvent,
		"NewTLVAdvancedCustomDrawEvent":        NewTLVAdvancedCustomDrawEvent,
		"NewTLVAdvancedCustomDrawItemEvent":    NewTLVAdvancedCustomDrawItemEvent,
		"NewTLVAdvancedCustomDrawSubItemEvent": NewTLVAdvancedCustomDrawSubItemEvent,
		"NewTLVChangeEvent":                    NewTLVChangeEvent,
		"NewTLVColumnClickEvent":               NewTLVColumnClickEvent,
		"NewTLVCompareEvent":                   NewTLVCompareEvent,
		"NewTLVOwnerDataEvent":                 NewTLVOwnerDataEvent,
		"NewTLVOwnerDataFindEvent":             NewTLVOwnerDataFindEvent,
		"NewTLVOwnerDataHintEvent":             NewTLVOwnerDataHintEvent,
		"NewTLVDataHintEvent":                  NewTLVDataHintEvent,
		"NewTLVDeletedEvent":                   NewTLVDeletedEvent,
		"NewTLVEditedEvent":                    NewTLVEditedEvent,
		"NewTLVEditingEvent":                   NewTLVEditingEvent,
		"NewTMenuChangeEvent":                  NewTMenuChangeEvent,
		"NewTMenuMeasureItemEvent":             NewTMenuMeasureItemEvent,
		"NewTTabChangingEvent":                 NewTTabChangingEvent,
		"NewTUDChangingEvent":                  NewTUDChangingEvent,
		"NewTUDClickEvent":                     NewTUDClickEvent,
		"NewTTaskDlgClickEvent":                NewTTaskDlgClickEvent,
		"NewTTVExpandedEvent":                  NewTTVExpandedEvent,
		"NewTTVAdvancedCustomDrawEvent":        NewTTVAdvancedCustomDrawEvent,
		"NewTTVAdvancedCustomDrawItemEvent":    NewTTVAdvancedCustomDrawItemEvent,
		"NewTTVChangedEvent":                   NewTTVChangedEvent,
		"NewTTVChangingEvent":                  NewTTVChangingEvent,
		"NewTTVCollapsingEvent":                NewTTVCollapsingEvent,
		"NewTTVCompareEvent":                   NewTTVCompareEvent,
		"NewTTVCustomDrawEvent":                NewTTVCustomDrawEvent,
		"NewTTVCustomDrawItemEvent":            NewTTVCustomDrawItemEvent,
		"NewTTVEditedEvent":                    NewTTVEditedEvent,
		"NewTTVEditingEvent":                   NewTTVEditingEvent,
		"NewTTVExpandingEvent":                 NewTTVExpandingEvent,

		"GetApplication": getVclApplication,
		// "NewApplication":    vcl.NewApplication,
		"InitVCL":           initLCL,
		"InitLCL":           initLCL,
		"NewCheckBox":       vcl.NewCheckBox,
		"NewLabel":          vcl.NewLabel,
		"NewButton":         vcl.NewButton,
		"NewComboBox":       vcl.NewComboBox,
		"NewEdit":           vcl.NewEdit,
		"NewCanvas":         vcl.NewCanvas,
		"NewImage":          vcl.NewImage,
		"NewList":           vcl.NewList,
		"NewListBox":        vcl.NewListBox,
		"NewListView":       vcl.NewListView,
		"NewListColumns":    vcl.NewListColumns,
		"NewListItem":       vcl.NewListItem,
		"NewListItems":      vcl.NewListItems,
		"NewMainMenu":       vcl.NewMainMenu,
		"NewMemo":           vcl.NewMemo,
		"NewMenuItem":       vcl.NewMenuItem,
		"NewMiniWebview":    vcl.NewMiniWebview,
		"NewPaintBox":       vcl.NewPaintBox,
		"NewPanel":          vcl.NewPanel,
		"NewPicture":        vcl.NewPicture,
		"NewPopupMenu":      vcl.NewPopupMenu,
		"NewProgressBar":    vcl.NewProgressBar,
		"NewRadioButton":    vcl.NewRadioButton,
		"NewRadioGroup":     vcl.NewRadioGroup,
		"NewScrollBox":      vcl.NewScrollBox,
		"NewScrollBar":      vcl.NewScrollBar,
		"NewSplitter":       vcl.NewSplitter,
		"NewStatusBar":      vcl.NewStatusBar,
		"NewStatusPanel":    vcl.NewStatusPanel,
		"NewStatusPanels":   vcl.NewStatusPanels,
		"NewTimer":          vcl.NewTimer,
		"NewToolBar":        vcl.NewToolBar,
		"NewToolButton":     vcl.NewToolButton,
		"NewTrayIcon":       vcl.NewTrayIcon,
		"NewStaticText":     vcl.NewStaticText,
		"NewSpinEdit":       vcl.NewSpinEdit,
		"NewSpeedButton":    vcl.NewSpeedButton,
		"NewShape":          vcl.NewShape,
		"NewScreen":         vcl.NewScreen,
		"NewSaveDialog":     vcl.NewSaveDialog,
		"NewReplaceDialog":  vcl.NewReplaceDialog,
		"NewPngImage":       vcl.NewPngImage,
		"NewPen":            vcl.NewPen,
		"NewPageControl":    vcl.NewPageControl,
		"NewOpenDialog":     vcl.NewOpenDialog,
		"NewObject":         vcl.NewObject,
		"NewMouse":          vcl.NewMouse,
		"NewMaskEdit":       vcl.NewMaskEdit,
		"NewLinkLabel":      vcl.NewLinkLabel,
		"NewLabeledEdit":    vcl.NewLabeledEdit,
		"NewJPEGImage":      vcl.NewJPEGImage,
		"NewImageList":      vcl.NewImageList,
		"NewImageButton":    vcl.NewImageButton,
		"NewIcon":           vcl.NewIcon,
		"NewGroupBox":       vcl.NewGroupBox,
		"NewHeaderControl":  vcl.NewHeaderControl,
		"NewHeaderSection":  vcl.NewHeaderSection,
		"NewHeaderSections": vcl.NewHeaderSections,
		"NewGraphic":        vcl.NewGraphic,
		"NewGIFImage":       vcl.NewGIFImage,
		"NewGauge":          vcl.NewGauge,
		"ShowMessage":       vcl.ShowMessage,
		"ShowMessageFmt":    vcl.ShowMessageFmt,
		"MessageDlg":        vcl.MessageDlg,
		"InputBox":          vcl.InputBox,
		"InputQuery":        vcl.InputQuery,
		"ThreadSync":        vcl.ThreadSync,
		"NewFrame":          vcl.NewFrame,
		"SelectDirectory":   vcl.SelectDirectory1,
		"SelectDirectory3":  vcl.SelectDirectory3,
		"NewForm":           vcl.NewForm,
		"NewFontDialog":     vcl.NewFontDialog,
		"NewFont":           vcl.NewFont,
		"NewFlowPanel":      vcl.NewFlowPanel,
		"NewFindDialog":     vcl.NewFindDialog,
		"NewDrawGrid":       vcl.NewDrawGrid,
		"NewDateTimePicker": vcl.NewDateTimePicker,
		"NewControl":        vcl.NewControl,
		"NewComboBoxEx":     vcl.NewComboBoxEx,
		"NewColorListBox":   vcl.NewColorListBox,
		"NewColorDialog":    vcl.NewColorDialog,
		"NewColorBox":       vcl.NewColorBox,
		"NewCheckListBox":   vcl.NewCheckListBox,
		"NewBrush":          vcl.NewBrush,
		"NewBitmap":         vcl.NewBitmap,
		"NewBitBtn":         vcl.NewBitBtn,
		"NewBevel":          vcl.NewBevel,
		"NewApplication":    vcl.NewApplication,
		"NewAction":         vcl.NewAction,
		"NewActionList":     vcl.NewActionList,
		"NewMemoryStream":   vcl.NewMemoryStream,

		// "NewAnchors": types.TAnchors,
		// "RTLInclude": rtl.Include,
		"NewSet":           types.NewSet,
		"AkTop":            types.AkTop,
		"AkBottom":         types.AkBottom,
		"AkLeft":           types.AkLeft,
		"AkRight":          types.AkRight,
		"SsNone":           types.SsNone,
		"SsHorizontal":     types.SsHorizontal,
		"SsVertical":       types.SsVertical,
		"SsBoth":           types.SsBoth,
		"SsAutoHorizontal": types.SsAutoHorizontal,
		"SsAutoVertical":   types.SsAutoVertical,
		"SsAutoBoth":       types.SsAutoBoth,

		"GetLibVersion": vcl.GetLibVersion,

		// values
		"PoDesigned":        types.PoDesigned,
		"PoDefault":         types.PoDefault,
		"PoDefaultPosOnly":  types.PoDefaultPosOnly,
		"PoDefaultSizeOnly": types.PoDefaultSizeOnly,
		"PoScreenCenter":    types.PoScreenCenter,
		"PoMainFormCenter":  types.PoMainFormCenter,
		"PoOwnerFormCenter": types.PoOwnerFormCenter,
		"PoWorkAreaCenter":  types.PoWorkAreaCenter,
	}

	qlang.Import("lcl", lclExports)

	qlang.Import("github_topxeq_govcl_vcl", qlgithub_topxeq_govcl_vcl.Exports)
	qlang.Import("vcl", qlgithub_topxeq_govcl_vcl.Exports)
	qlang.Import("github_topxeq_govcl_vcl_types", qlgithub_topxeq_govcl_vcl_types.Exports)
	qlang.Import("vcl_types", qlgithub_topxeq_govcl_vcl_types.Exports)
	qlang.Import("github_topxeq_govcl_vcl_api", qlgithub_topxeq_govcl_vcl_api.Exports)
	qlang.Import("vcl_api", qlgithub_topxeq_govcl_vcl_api.Exports)
	qlang.Import("github_topxeq_govcl_vcl_rtl", qlgithub_topxeq_govcl_vcl_rtl.Exports)
	qlang.Import("vcl_rtl", qlgithub_topxeq_govcl_vcl_rtl.Exports)

}

func initLCLLib() (result error) {

	defer func() {
		if r := recover(); r != nil {
			// fmt.Printf("initLCLLib: %v\n", r)

			result = tk.Errf("initLCLLib: %v\n", r)
		}
	}()

	api.DoLibInit()

	result = nil

	return result
}

// func syncInitLCL() error {
// 	var errT error

// 	vcl.ThreadSync(func() {
// 		initLCL()
// 	})

// 	return errT
// }

func initLCL() error {

	runtime.LockOSThread()

	// startThreadID := tk.GetCurrentThreadID()

	api.CloseLib()

	errT := initLCLLib()

	if errT != nil {
		tk.Pl("failed to init lib: %v, try to download the LCL lib...", errT)

		applicationPathT := tk.GetApplicationPath()

		osT := tk.GetOSName()

		if tk.Contains(osT, "inux") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.so", applicationPathT, "liblcl.so", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		} else if tk.Contains(osT, "arwin") {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dylib", applicationPathT, "liblcl.dylib", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		} else {
			rs := tk.DownloadFile("http://scripts.frenchfriend.net/pub/liblcl.dll", applicationPathT, "liblcl.dll", false)

			if tk.IsErrorString(rs) {
				tk.Pl("failed to download LCL file.")
				return tk.Errf("failed to download LCL file.")
			}
		}

		errT = initLCLLib()

		if errT != nil {
			tk.Pl("failed to install lib: %v", errT)
			return tk.Errf("failed to install lib: %v", errT)
		}
	}

	api.DoResInit()

	api.DoImportInit()

	api.DoDefInit()

	rtl.DoRtlInit()

	vcl.DoInit()

	// if verboseG {

	// 	endThreadID := tk.GetCurrentThreadID()

	// 	tk.Pl("start tid: %v, end tid: %v", startThreadID, endThreadID)

	// 	if endThreadID != startThreadID {
	// 		return tk.Errf("failed to init lcl lib: %v", "thread not same")
	// 	}
	// }

	return nil
}

func getVclApplication() *vcl.TApplication {
	return vcl.Application
}

// func newLclAnchors() *types.TAnchors {
// 	a := types.TAnchors(rtl.Include(0, types.AkTop, types.AkBottom, types.AkLeft, types.AkRight))

// }

// func getLCLEvent(funcA *execq.Function) *() {

// }

// func getTNotifyEvent(funcA *execq.Function) *vcl.TNotifyEvent {
// 	var f vcl.TNotifyEvent = func(sender vcl.IObject) {
// 		funcA.Call(execq.NewStack(), sender)
// 	}

// 	return &f
// }

func NewTNotifyEvent(funcA *execq.Function) *vcl.TNotifyEvent {
	var f vcl.TNotifyEvent = func(sender vcl.IObject) {
		funcA.Call(execq.NewStack(), sender)
	}

	return &f
}

func NewTKeyEvent(funcA *execq.Function) *vcl.TKeyEvent {
	var f vcl.TKeyEvent = func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		funcA.Call(execq.NewStack(), sender, key, shift)
	}

	return &f
	// return nil
}

func NewTKeyPressEvent(funcA *execq.Function) *vcl.TKeyPressEvent {
	var f vcl.TKeyPressEvent = func(sender vcl.IObject, key *types.Char) {
		funcA.Call(execq.NewStack(), sender, key)
	}

	return &f
	// return nil
}

func NewTMouseEvent(funcA *execq.Function) *vcl.TMouseEvent {
	var f vcl.TMouseEvent = func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		funcA.Call(execq.NewStack(), sender, button, shift, x, y)
	}

	return &f
	// return nil
}

func NewTMouseMoveEvent(funcA *execq.Function) *vcl.TMouseMoveEvent {
	var f vcl.TMouseMoveEvent = func(sender vcl.IObject, shift types.TShiftState, x, y int32) {
		funcA.Call(execq.NewStack(), sender, shift, x, y)
	}

	return &f
	// return nil
}

func NewTExceptionEvent(funcA *execq.Function) *vcl.TExceptionEvent {
	var f vcl.TExceptionEvent = func(sender vcl.IObject, e *vcl.Exception) {
		funcA.Call(execq.NewStack(), sender, e)
	}

	return &f
}

func NewTCloseEvent(funcA *execq.Function) *vcl.TCloseEvent {
	var f vcl.TCloseEvent = func(sender vcl.IObject, action *types.TCloseAction) { // TCloseAction int32
		funcA.Call(execq.NewStack(), sender, action)
	}

	return &f
}

func NewTCloseQueryEvent(funcA *execq.Function) *vcl.TCloseQueryEvent {
	var f vcl.TCloseQueryEvent = func(sender vcl.IObject, canClose *bool) {
		funcA.Call(execq.NewStack(), sender, canClose)
	}

	return &f
}

func NewTContextPopupEvent(funcA *execq.Function) *vcl.TContextPopupEvent {
	var f vcl.TContextPopupEvent = func(sender vcl.IObject, mousePos types.TPoint, handled *bool) {
		funcA.Call(execq.NewStack(), sender, mousePos, handled)
	}

	return &f
}

func NewTDragDropEvent(funcA *execq.Function) *vcl.TDragDropEvent {
	var f vcl.TDragDropEvent = func(sender vcl.IObject, source vcl.IObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, source, x, y)
	}

	return &f
}

func NewTDragOverEvent(funcA *execq.Function) *vcl.TDragOverEvent {
	var f vcl.TDragOverEvent = func(sender vcl.IObject, source vcl.IObject, x, y int32, state types.TDragState, accept *bool) {
		funcA.Call(execq.NewStack(), sender, source, x, y, state, accept)
	}

	return &f
}

func NewTStartDragEvent(funcA *execq.Function) *vcl.TStartDragEvent {
	var f vcl.TStartDragEvent = func(sender vcl.IObject, dragObject *vcl.TDragObject) {
		funcA.Call(execq.NewStack(), sender, dragObject)
	}

	return &f
}

func NewTEndDragEvent(funcA *execq.Function) *vcl.TEndDragEvent {
	var f vcl.TEndDragEvent = func(sender vcl.IObject, target vcl.IObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, target, x, y)
	}

	return &f
}

func NewTAlignPositionEvent(funcA *execq.Function) *vcl.TAlignPositionEvent {
	var f vcl.TAlignPositionEvent = func(sender *vcl.TWinControl, control *vcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
		funcA.Call(execq.NewStack(), sender, control, newLeft, newTop, newWidth, newHeight, alignRect, alignInfo)
	}

	return &f
}

func NewTDockDropEvent(funcA *execq.Function) *vcl.TDockDropEvent {
	var f vcl.TDockDropEvent = func(sender vcl.IObject, source *vcl.TDragDockObject, x, y int32) {
		funcA.Call(execq.NewStack(), sender, source, x, y)
	}

	return &f
}

func NewTDockOverEvent(funcA *execq.Function) *vcl.TDockOverEvent {
	var f vcl.TDockOverEvent = func(sender vcl.IObject, source *vcl.TDragDockObject, x, y int32, state types.TDragState, accept *bool) {
		funcA.Call(execq.NewStack(), sender, source, x, y, state, accept)
	}

	return &f
}

func NewTStartDockEvent(funcA *execq.Function) *vcl.TStartDockEvent {
	var f vcl.TStartDockEvent = func(sender vcl.IObject, dragObject *vcl.TDragDockObject) {
		funcA.Call(execq.NewStack(), sender, dragObject)
	}

	return &f
}

func NewTUnDockEvent(funcA *execq.Function) *vcl.TUnDockEvent {
	var f vcl.TUnDockEvent = func(sender vcl.IObject, client *vcl.TControl, newTarget *vcl.TControl, allow *bool) {
		funcA.Call(execq.NewStack(), sender, client, newTarget, allow)
	}

	return &f
}

func NewTGetSiteInfoEvent(funcA *execq.Function) *vcl.TGetSiteInfoEvent {
	var f vcl.TGetSiteInfoEvent = func(sender vcl.IObject, dockClient *vcl.TControl, influenceRect *types.TRect, mousePos types.TPoint, canDock *bool) {
		funcA.Call(execq.NewStack(), sender, dockClient, influenceRect, mousePos, canDock)
	}

	return &f
}

func NewTMouseWheelEvent(funcA *execq.Function) *vcl.TMouseWheelEvent {
	var f vcl.TMouseWheelEvent = func(sender vcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
		funcA.Call(execq.NewStack(), sender, shift, wheelDelta, x, y, handled)
	}

	return &f
}

func NewTMouseWheelUpDownEvent(funcA *execq.Function) *vcl.TMouseWheelUpDownEvent {
	var f vcl.TMouseWheelUpDownEvent = func(sender vcl.IObject, shift types.TShiftState, mousePos types.TPoint, handled *bool) {
		funcA.Call(execq.NewStack(), sender, shift, mousePos, handled)
	}

	return &f
}

func NewTMessageEvent(funcA *execq.Function) *vcl.TMessageEvent {
	var f vcl.TMessageEvent = func(msg *types.TMsg, handled *bool) {
		funcA.Call(execq.NewStack(), msg, handled)
	}

	return &f
}

func NewTHelpEvent(funcA *execq.Function) *vcl.THelpEvent {
	var f vcl.THelpEvent = func(command uint16, data types.THelpEventData, callhelp, result *bool) {
		funcA.Call(execq.NewStack(), command, data, callhelp, result)
	}

	return &f
}

func NewTWebTitleChangeEvent(funcA *execq.Function) *vcl.TWebTitleChangeEvent {
	var f vcl.TWebTitleChangeEvent = func(sender vcl.IObject, text string) {
		funcA.Call(execq.NewStack(), sender, text)
	}

	return &f
}

func NewTWebJSExternalEvent(funcA *execq.Function) *vcl.TWebJSExternalEvent {
	var f vcl.TWebJSExternalEvent = func(sender vcl.IObject, funcName, args string, retVal *string) {
		funcA.Call(execq.NewStack(), sender, funcName, args, retVal)
	}

	return &f
}

func NewTMeasureItemEvent(funcA *execq.Function) *vcl.TMeasureItemEvent {
	var f vcl.TMeasureItemEvent = func(control *vcl.TWinControl, index int32, height *int32) {
		funcA.Call(execq.NewStack(), control, index, height)
	}

	return &f
}

func NewTMovedEvent(funcA *execq.Function) *vcl.TMovedEvent {
	var f vcl.TMovedEvent = func(sender vcl.IObject, fromIndex, toIndex int32) {
		funcA.Call(execq.NewStack(), sender, fromIndex, toIndex)
	}

	return &f
}

func NewTDrawCellEvent(funcA *execq.Function) *vcl.TDrawCellEvent {
	var f vcl.TDrawCellEvent = func(sender vcl.IObject, aCol, aRow int32, aRect types.TRect, state types.TGridDrawState) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, aRect, state)
	}

	return &f
}

func NewTSelectCellEvent(funcA *execq.Function) *vcl.TSelectCellEvent {
	var f vcl.TSelectCellEvent = func(sender vcl.IObject, aCol, aRow int32, canSelect *bool) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, canSelect)
	}

	return &f
}

func NewTGetEditEvent(funcA *execq.Function) *vcl.TGetEditEvent {
	var f vcl.TGetEditEvent = func(sender vcl.IObject, aCol, aRow int32, value *string) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, value)
	}

	return &f
}

func NewTSetEditEvent(funcA *execq.Function) *vcl.TSetEditEvent {
	var f vcl.TSetEditEvent = func(sender vcl.IObject, aCol, aRow int32, value string) {
		funcA.Call(execq.NewStack(), sender, aCol, aRow, value)
	}

	return &f
}

func NewTDropFilesEvent(funcA *execq.Function) *vcl.TDropFilesEvent {
	var f vcl.TDropFilesEvent = func(sender vcl.IObject, aFileNames []string) {
		funcA.Call(execq.NewStack(), sender, aFileNames)
	}

	return &f
}

func NewTConstrainedResizeEvent(funcA *execq.Function) *vcl.TConstrainedResizeEvent {
	var f vcl.TConstrainedResizeEvent = func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		funcA.Call(execq.NewStack(), sender, minWidth, minHeight, maxWidth, maxHeight)
	}

	return &f
}

func NewTWndProcEvent(funcA *execq.Function) *vcl.TWndProcEvent {
	var f vcl.TWndProcEvent = func(msg *types.TMessage) {
		funcA.Call(execq.NewStack(), msg)
	}

	return &f
}

func NewTSectionNotifyEvent(funcA *execq.Function) *vcl.TSectionNotifyEvent {
	var f vcl.TSectionNotifyEvent = func(headerControl *vcl.THeaderControl, section *vcl.THeaderSection) {
		funcA.Call(execq.NewStack(), headerControl, section)
	}

	return &f
}

func NewTSectionTrackEvent(funcA *execq.Function) *vcl.TSectionTrackEvent {
	var f vcl.TSectionTrackEvent = func(headerControl *vcl.THeaderControl, section *vcl.THeaderSection, width int32, state types.TSectionTrackState) {
		funcA.Call(execq.NewStack(), headerControl, section, width, state)
	}

	return &f
}

func NewTSectionDragEvent(funcA *execq.Function) *vcl.TSectionDragEvent {
	var f vcl.TSectionDragEvent = func(sender vcl.IObject, fromSection, toSection *vcl.THeaderSection, allowDrag *bool) {
		funcA.Call(execq.NewStack(), sender, fromSection, toSection, allowDrag)
	}

	return &f
}

func NewTSysLinkEvent(funcA *execq.Function) *vcl.TSysLinkEvent {
	var f vcl.TSysLinkEvent = func(sender vcl.IObject, link string, linkType types.TSysLinkType) {
		funcA.Call(execq.NewStack(), sender, link, linkType)
	}

	return &f
}

func NewTDrawItemEvent(funcA *execq.Function) *vcl.TDrawItemEvent {
	var f vcl.TDrawItemEvent = func(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
		funcA.Call(execq.NewStack(), control, index, aRect, state)
	}

	return &f
}

func NewTLVSelectItemEvent(funcA *execq.Function) *vcl.TLVSelectItemEvent {
	var f vcl.TLVSelectItemEvent = func(sender vcl.IObject, item *vcl.TListItem, selected bool) {
		funcA.Call(execq.NewStack(), sender, item, selected)
	}

	return &f
}

func NewTLVCheckedItemEvent(funcA *execq.Function) *vcl.TLVCheckedItemEvent {
	var f vcl.TLVCheckedItemEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVAdvancedCustomDrawEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawEvent {
	var f vcl.TLVAdvancedCustomDrawEvent = func(sender *vcl.TListView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, stage, defaultDraw)
	}

	return &f
}

func NewTLVAdvancedCustomDrawItemEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawItemEvent {
	var f vcl.TLVAdvancedCustomDrawItemEvent = func(sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, item, state, Stage, defaultDraw)
	}

	return &f
}

func NewTLVAdvancedCustomDrawSubItemEvent(funcA *execq.Function) *vcl.TLVAdvancedCustomDrawSubItemEvent {
	var f vcl.TLVAdvancedCustomDrawSubItemEvent = func(sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, item, subItem, state, stage, defaultDraw)
	}

	return &f
}

func NewTLVChangeEvent(funcA *execq.Function) *vcl.TLVChangeEvent {
	var f vcl.TLVChangeEvent = func(sender vcl.IObject, item *vcl.TListItem, change types.TItemChange) {
		funcA.Call(execq.NewStack(), sender, item, change)
	}

	return &f
}

func NewTLVColumnClickEvent(funcA *execq.Function) *vcl.TLVColumnClickEvent {
	var f vcl.TLVColumnClickEvent = func(sender vcl.IObject, column *vcl.TListColumn) {
		funcA.Call(execq.NewStack(), sender, column)
	}

	return &f
}

func NewTLVCompareEvent(funcA *execq.Function) *vcl.TLVCompareEvent {
	var f vcl.TLVCompareEvent = func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
		funcA.Call(execq.NewStack(), sender, item1, item2, data, compare)
	}

	return &f
}

func NewTLVOwnerDataEvent(funcA *execq.Function) *vcl.TLVOwnerDataEvent {
	var f vcl.TLVOwnerDataEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVOwnerDataFindEvent(funcA *execq.Function) *vcl.TLVOwnerDataFindEvent {
	var f vcl.TLVOwnerDataFindEvent = func(sender vcl.IObject, find types.TItemFind, findString string, findPosition types.TPoint, findData types.TCustomData, startIndex int32, direction types.TSearchDirection, warp bool, index *int32) {
		funcA.Call(execq.NewStack(), sender, find, findString, findPosition, findData, startIndex, direction, warp, index)
	}

	return &f
}

func NewTLVOwnerDataHintEvent(funcA *execq.Function) *vcl.TLVOwnerDataHintEvent {
	var f vcl.TLVOwnerDataHintEvent = func(sender vcl.IObject, startIndex, endIndex int32) {
		funcA.Call(execq.NewStack(), sender, startIndex, endIndex)
	}

	return &f
}

func NewTLVDataHintEvent(funcA *execq.Function) *vcl.TLVDataHintEvent {
	var f vcl.TLVDataHintEvent = func(sender vcl.IObject, startIndex, endIndex int32) {
		funcA.Call(execq.NewStack(), sender, startIndex, endIndex)
	}

	return &f
}

func NewTLVDeletedEvent(funcA *execq.Function) *vcl.TLVDeletedEvent {
	var f vcl.TLVDeletedEvent = func(sender vcl.IObject, item *vcl.TListItem) {
		funcA.Call(execq.NewStack(), sender, item)
	}

	return &f
}

func NewTLVEditedEvent(funcA *execq.Function) *vcl.TLVEditedEvent {
	var f vcl.TLVEditedEvent = func(sender vcl.IObject, item *vcl.TListItem, s *string) {
		funcA.Call(execq.NewStack(), sender, item, s)
	}

	return &f
}

func NewTLVEditingEvent(funcA *execq.Function) *vcl.TLVEditingEvent {
	var f vcl.TLVEditingEvent = func(sender vcl.IObject, item *vcl.TListItem, allowEdit *bool) {
		funcA.Call(execq.NewStack(), sender, item, allowEdit)
	}

	return &f
}

func NewTMenuChangeEvent(funcA *execq.Function) *vcl.TMenuChangeEvent {
	var f vcl.TMenuChangeEvent = func(sender vcl.IObject, source *vcl.TMenuItem, rebuild bool) {
		funcA.Call(execq.NewStack(), sender, source, rebuild)
	}

	return &f
}

func NewTMenuMeasureItemEvent(funcA *execq.Function) *vcl.TMenuMeasureItemEvent {
	var f vcl.TMenuMeasureItemEvent = func(sender vcl.IObject, aCanvas *vcl.TCanvas, width, height *int32) {
		funcA.Call(execq.NewStack(), sender, aCanvas, width, height)
	}

	return &f
}

func NewTTabChangingEvent(funcA *execq.Function) *vcl.TTabChangingEvent {
	var f vcl.TTabChangingEvent = func(sender vcl.IObject, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, allowChange)
	}

	return &f
}

func NewTUDChangingEvent(funcA *execq.Function) *vcl.TUDChangingEvent {
	var f vcl.TUDChangingEvent = func(sender vcl.IObject, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, allowChange)
	}

	return &f
}

func NewTUDClickEvent(funcA *execq.Function) *vcl.TUDClickEvent {
	var f vcl.TUDClickEvent = func(sender vcl.IObject, button types.TUDBtnType) {
		funcA.Call(execq.NewStack(), sender, button)
	}

	return &f
}

func NewTTaskDlgClickEvent(funcA *execq.Function) *vcl.TTaskDlgClickEvent {
	var f vcl.TTaskDlgClickEvent = func(sender vcl.IObject, modalResult types.TModalResult, canClose *bool) {
		funcA.Call(execq.NewStack(), sender, modalResult, canClose)
	}

	return &f
}

func NewTTVExpandedEvent(funcA *execq.Function) *vcl.TTVExpandedEvent {
	var f vcl.TTVExpandedEvent = func(sender vcl.IObject, node *vcl.TTreeNode) {
		funcA.Call(execq.NewStack(), sender, node)
	}

	return &f
}

func NewTTVAdvancedCustomDrawEvent(funcA *execq.Function) *vcl.TTVAdvancedCustomDrawEvent {
	var f vcl.TTVAdvancedCustomDrawEvent = func(sender *vcl.TTreeView, aRect types.TRect, stage types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, stage, defaultDraw)
	}

	return &f
}

func NewTTVAdvancedCustomDrawItemEvent(funcA *execq.Function) *vcl.TTVAdvancedCustomDrawItemEvent {
	var f vcl.TTVAdvancedCustomDrawItemEvent = func(sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawState, stage types.TCustomDrawStage, paintImages, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, node, state, stage, paintImages, defaultDraw)
	}

	return &f
}

func NewTTVChangedEvent(funcA *execq.Function) *vcl.TTVChangedEvent {
	var f vcl.TTVChangedEvent = func(sender vcl.IObject, node *vcl.TTreeNode) {
		funcA.Call(execq.NewStack(), sender, node)
	}

	return &f
}

func NewTTVChangingEvent(funcA *execq.Function) *vcl.TTVChangingEvent {
	var f vcl.TTVChangingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowChange *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowChange)
	}

	return &f
}

func NewTTVCollapsingEvent(funcA *execq.Function) *vcl.TTVCollapsingEvent {
	var f vcl.TTVCollapsingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowCollapse *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowCollapse)
	}

	return &f
}

func NewTTVCompareEvent(funcA *execq.Function) *vcl.TTVCompareEvent {
	var f vcl.TTVCompareEvent = func(sender vcl.IObject, node1, node2 *vcl.TTreeNode, data int32, compare *int32) {
		funcA.Call(execq.NewStack(), sender, node1, node2, data, compare)
	}

	return &f
}

func NewTTVCustomDrawEvent(funcA *execq.Function) *vcl.TTVCustomDrawEvent {
	var f vcl.TTVCustomDrawEvent = func(sender *vcl.TTreeView, aRect types.TRect, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, aRect, defaultDraw)
	}

	return &f
}

func NewTTVCustomDrawItemEvent(funcA *execq.Function) *vcl.TTVCustomDrawItemEvent {
	var f vcl.TTVCustomDrawItemEvent = func(sender *vcl.TTreeView, node *vcl.TTreeNode, state types.TCustomDrawStage, defaultDraw *bool) {
		funcA.Call(execq.NewStack(), sender, node, state, defaultDraw)
	}

	return &f
}

func NewTTVEditedEvent(funcA *execq.Function) *vcl.TTVEditedEvent {
	var f vcl.TTVEditedEvent = func(sender vcl.IObject, node *vcl.TTreeNode, s *string) {
		funcA.Call(execq.NewStack(), sender, node, s)
	}

	return &f
}

func NewTTVEditingEvent(funcA *execq.Function) *vcl.TTVEditingEvent {
	var f vcl.TTVEditingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowEdit *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowEdit)
	}

	return &f
}

func NewTTVExpandingEvent(funcA *execq.Function) *vcl.TTVExpandingEvent {
	var f vcl.TTVExpandingEvent = func(sender vcl.IObject, node *vcl.TTreeNode, allowExpansion *bool) {
		funcA.Call(execq.NewStack(), sender, node, allowExpansion)
	}

	return &f
}
