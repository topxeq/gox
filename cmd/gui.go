//go:build !linux && !darwin
// +build !linux,!darwin

package main

import (
	"fmt"
	"runtime"

	jsoniter "github.com/json-iterator/go"
	"github.com/ncruces/zenity"
	"github.com/topxeq/dlgs"
	"github.com/topxeq/xie"

	"github.com/topxeq/tk"

	"github.com/kbinani/screenshot"

	"github.com/jchv/go-webview2"
)

// ...
// "github.com/jchv/go-webview2"
// "github.com/jchv/go-webview2/pkg/edge"
// )

// func main() {
// dataPath, _ := filepath.Abs("./userdata")
// w := webview2.NewWithOptions(webview2.WebViewOptions{
// 	Debug:     true,
// 	AutoFocus: true,
// 	DataPath:  dataPath,
// 	WindowOptions: webview2.WindowOptions{
// 		Title: "go-webview2 Example",
// 	},
// })
// if w == nil {
// 	log.Fatalln("Failed to load webview.")
// }
// defer w.Destroy()

// // update window icon
// w32.SendMessage(w.Window(), 0x0080, 1, w32.ExtractIcon(os.Args[0], 0))

// w.SetSize(800, 600, webview2.HintNone)

// chromium := getChromium(w)

// folderPath, _ := filepath.Abs("./public")
// webview := chromium.GetICoreWebView2_3()
// webview.SetVirtualHostNameToFolderMapping(
// 	"app.assets", folderPath,
// 	edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY_CORS,
// )
// w.Navigate("http://app.assets/index.html")

// w.Run()
// }

// func getChromium(w webview2.WebView) *edge.Chromium {
// browser := reflect.ValueOf(w).Elem().FieldByName("browser")
// browser = reflect.NewAt(browser.Type(), unsafe.Pointer(browser.UnsafeAddr())).Elem()
// return browser.Interface().(*edge.Chromium)
// }

var windowStyleG = webview2.HintNone

func newWindowWebView2(objA interface{}, paramsA []interface{}) interface{} {
	var paraArgsT []string = []string{}

	for i := 0; i < len(paramsA); i++ {
		paraArgsT = append(paraArgsT, tk.ToStr(paramsA[i]))
	}

	p := objA.(*xie.XieVM)

	titleT := p.GetSwitchVarValue(p.Running, paraArgsT, "-title=", "dialog")
	widthT := p.GetSwitchVarValue(p.Running, paraArgsT, "-width=", "800")
	heightT := p.GetSwitchVarValue(p.Running, paraArgsT, "-height=", "600")
	iconT := p.GetSwitchVarValue(p.Running, paraArgsT, "-icon=", "2")
	debugT := tk.IfSwitchExistsWhole(paraArgsT, "-debug")
	centerT := tk.IfSwitchExistsWhole(paraArgsT, "-center")
	fixT := tk.IfSwitchExistsWhole(paraArgsT, "-fix")
	maxT := tk.IfSwitchExistsWhole(paraArgsT, "-max")
	minT := tk.IfSwitchExistsWhole(paraArgsT, "-min")

	if maxT {
		// windowStyleT = webview2.HintMax

		rectT := screenshot.GetDisplayBounds(0)

		widthT = tk.ToStr(rectT.Max.X)
		heightT = tk.ToStr(rectT.Max.Y)
	}

	if minT {
		// windowStyleT = webview2.HintMin

		widthT = "0"
		heightT = "0"
	}

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     debugT,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  titleT,
			Width:  uint(tk.ToInt(widthT, 800)),
			Height: uint(tk.ToInt(heightT, 600)),
			IconId: uint(tk.ToInt(iconT, 2)), // icon resource id
			Center: centerT,
		},
	})

	if w == nil {
		return fmt.Errorf("创建窗口失败：%v", "N/A")
	}

	windowStyleG = webview2.HintNone

	if fixT {
		windowStyleG = webview2.HintFixed
	}

	w.SetSize(tk.ToInt(widthT, 800), tk.ToInt(heightT, 600), windowStyleG)

	w.Bind("delegateCloseWindow", func(argsA ...interface{}) interface{} {
		w.Destroy()

		return ""
	})

	var handlerT tk.TXDelegate

	handlerT = func(actionA string, objA interface{}, dataA interface{}, paramsA ...interface{}) interface{} {
		switch actionA {
		case "show":
			w.Run()
			return nil
		case "setSize":
			len1T := len(paramsA)
			if len1T < 2 {
				return fmt.Errorf("参数不够")
			}
			if len1T > 2 {
				windowStyleG = webview2.Hint(tk.ToInt(paramsA[2]))
			}

			w.SetSize(tk.ToInt(paramsA[0], 800), tk.ToInt(paramsA[1], 600), windowStyleG)
			return nil
		case "navigate":
			len1T := len(paramsA)
			if len1T < 1 {
				return fmt.Errorf("参数不够")
			}

			if len1T > 0 {
				w.Navigate(tk.ToStr(paramsA[0]))
			}

			return nil
		case "setHtml":
			len1T := len(paramsA)
			if len1T < 1 {
				return fmt.Errorf("参数不够")
			}

			if len1T > 0 {
				w.SetHtml(tk.ToStr(paramsA[0]))
			}

			return nil
		case "call", "eval":
			len1T := len(paramsA)
			if len1T < 1 {
				return fmt.Errorf("参数不够")
			}

			if len1T > 0 {
				w.Dispatch(func() {
					w.Eval(tk.ToStr(paramsA[0]))
				})
			}

			return nil
		case "close":
			w.Destroy()
			return nil
		case "setQuickDelegate":
			len1T := len(paramsA)
			if len1T < 1 {
				return fmt.Errorf("not enough parameters")
			}

			var deleT tk.QuickVarDelegate
			var ok bool
			var s1 string

			codeT := paramsA[0]

			deleT, ok = codeT.(tk.QuickVarDelegate)
			if ok {

			} else {
				if s1, ok = codeT.(string); ok {
					// s1 = strings.ReplaceAll(s1, "~~~", "`")
					compiledT := xie.Compile(s1)

					if tk.IsError(compiledT) {
						return fmt.Errorf("failed to compile the quick delegate code: %v", compiledT)
					}

					codeT = compiledT
				}

				cp1, ok := codeT.(*xie.CompiledCode)

				if ok {
					deleT = func(argsA ...interface{}) interface{} {
						rs := xie.RunCodePiece(p, nil, cp1, argsA, true)

						return rs
					}
				} else {
					return fmt.Errorf("invalid compiled object: %v", codeT)
				}

			}

			// deleT, ok := paramsA[0].(tk.QuickVarDelegate)

			// if !ok {
			// 	s1, ok := paramsA[0].(string)

			// 	if ！ok {
			// 		return fmt.Errorf("unknown type: %T(%v)", paramsA[0], paramsA[0])
			// 	}
			// }

			w.Bind("quickDelegateDo", func(args ...interface{}) interface{} {
				// args是WebView2中调用谢语言函数时传入的参数
				// 可以是多个，谢语言中按位置索引进行访问
				// strT := args[0].String()

				rsT := deleT(args...)

				if tk.IsErrX(rsT) {
					if xie.GlobalsG.VerboseLevel > 0 {
						tk.Pl("error occurred in QuickVarDelegate: %v", rsT)
					}
				}

				// 最后一定要返回一个值，空字符串也可以
				return rsT
			})

			return nil
		case "setDelegate":
			len1T := len(paramsA)
			if len1T < 1 {
				return fmt.Errorf("not enough parameters")
			}
			// tk.Pl("----")
			// tk.Plo(paramsA)
			// tk.Pl("----")

			var codeT = paramsA[0]

			var deleT tk.QuickVarDelegate
			var ok bool
			var s1 string

			deleT, ok = codeT.(tk.QuickVarDelegate)
			if ok {
				errT := w.Bind("delegateDo", deleT)
				return errT
			} else {
				if s1, ok = codeT.(string); ok {
					// s1 = strings.ReplaceAll(s1, "~~~", "`")
					compiledT := xie.Compile(s1)

					if tk.IsError(compiledT) {
						return fmt.Errorf("failed to compile the quick delegate code: %v", compiledT)
					}

					codeT = compiledT
				}

				cp1, ok := codeT.(*xie.CompiledCode)

				if ok {
					vmT := xie.NewVMQuick()

					lrs := vmT.Load(nil, cp1)

					if tk.IsError(lrs) {
						return fmt.Errorf("failed to create VM: %v", lrs)
					}

					deleT = func(argsA ...interface{}) interface{} {
						vmT.SetVar(nil, "inputG", argsA)

						rs := vmT.Run()

						return rs
					}
				} else {
					return fmt.Errorf("invalid compiled object: %v", codeT)
				}

			}

			// deleT, ok := paramsA[0].(tk.QuickVarDelegate)

			// if !ok {
			// 	s1, ok := paramsA[0].(string)

			// 	if ！ok {
			// 		return fmt.Errorf("unknown type: %T(%v)", paramsA[0], paramsA[0])
			// 	}
			// }

			// tk.Pl("here: %#v", deleT)
			errT := w.Bind("delegateDo", deleT)
			// errT := w.Bind("delegateDo", func(args ...interface{}) interface{} {
			// 	// args是WebView2中调用谢语言函数时传入的参数
			// 	// 可以是多个，谢语言中按位置索引进行访问
			// 	// strT := args[0].String()

			// 	rsT := deleT(args...)

			// 	if tk.IsErrX(rsT) {
			// 		if xie.GlobalsG.VerboseLevel > 0 {
			// 			tk.Pl("error occurred in QuickVarDelegate: %v", rsT)
			// 		}
			// 	}

			// 	// 最后一定要返回一个值，空字符串也可以
			// 	return rsT
			// })

			return errT

			// dv1, ok := codeT.(tk.QuickVarDelegate)

			// if ok {
			// 	w.Bind("delegateDo", dv1)
			// 	return nil
			// }

			// sv1, ok := codeT.(string)

			// if ok {
			// 	sv1 = strings.ReplaceAll(sv1, "~~~", "`")

			// 	vmT := xie.NewVMQuick()

			// 	vmT.SetVar(p.Running, "guiG", guiHandler)

			// 	lrs := vmT.Load(nil, sv1) // vmT.Running

			// 	if tk.IsError(lrs) {
			// 		return lrs
			// 	}

			// 	w.Bind("delegateDo", func(argsA ...interface{}) interface{} {
			// 		// args是WebView2中调用谢语言函数时传入的参数
			// 		// 可以是多个，谢语言中按位置索引进行访问
			// 		// strT := args[0].String()

			// 		vmT.SetVar(nil, "inputG", argsA) // p.Running

			// 		rs := vmT.Run()

			// 		// if !tk.IsErrX(rs) {
			// 		// 	outIndexT, ok := vmT.VarIndexMapM["outG"]
			// 		// 	if !ok {
			// 		// 		return tk.ErrStrf("no result")
			// 		// 	}

			// 		// 	return tk.ToStr((*vmT.FuncContextM.VarsM)[vmT.FuncContextM.VarsLocalMapM[outIndexT]])
			// 		// }

			// 		// 最后一定要返回一个值，空字符串也可以
			// 		return rs
			// 	})

			// }

			// p := objA.(*xie.XieVM)

			// return fmt.Errorf("invalid type: %T(%v)", codeT, codeT)
			// return nil
		case "setGoDelegate":
			var codeT string = tk.ToStr(paramsA[0])

			// p := objA.(*xie.XieVM)

			w.Bind("goDelegateDo", func(args ...interface{}) interface{} {
				// args是WebView2中调用谢语言函数时传入的参数
				// 可以是多个，谢语言中按位置索引进行访问
				// strT := args[0].String()

				vmT := xie.NewVMQuick()

				// xie.GlobalsG.Vars["verbose"]. = p.VerboseM
				// vmT.VerbosePlusM = p.VerbosePlusM

				vmT.SetVar(vmT.Running, "inputG", args)

				// argCountT := p.Pop()

				// if argCountT == Undefined {
				// 	return tk.ErrStrf()
				// }

				// for i := 0; i < argCountA; i++ {
				// 	vmT.Push(p.Pop())
				// }

				lrs := vmT.Load(vmT.Running, codeT)

				if tk.IsErrX(lrs) {
					return lrs
				}

				go vmT.Run()

				// 最后一定要返回一个值，空字符串也可以
				return ""
			})

			return nil
		default:
			return fmt.Errorf("未知操作：%v", actionA)
		}

		return nil
	}

	// w.Show()
	// w.Run()

	return handlerT

}

func guiHandler(actionA string, objA interface{}, dataA interface{}, paramsA ...interface{}) interface{} {
	switch actionA {
	case "init":
		rs := initGUI()
		return rs
	case "lockOSThread":
		runtime.LockOSThread()
		return nil
	case "method", "mt":
		if len(paramsA) < 1 {
			return fmt.Errorf("参数不够")
		}

		objT := paramsA[0]

		methodNameT := tk.ToStr(paramsA[1])

		v1p := 2

		switch nv := objT.(type) {
		case zenity.ProgressDialog:
			switch methodNameT {
			case "close":
				rs := nv.Close()
				return rs
			case "complete":
				rs := nv.Complete()
				return rs
			case "text":
				if len(paramsA) < v1p+1 {
					return fmt.Errorf("参数不够")
				}

				v1 := tk.ToStr(paramsA[v1p])

				rs := nv.Text(v1)
				return rs
			case "value":
				if len(paramsA) < v1p+1 {
					return fmt.Errorf("参数不够")
				}

				v1 := tk.ToInt(paramsA[v1p])

				rs := nv.Value(v1)
				return rs
			case "maxValue":
				rs := nv.MaxValue()
				return rs
			case "done":
				return nv.Done()
			}
		}

		rvr := tk.ReflectCallMethod(objT, methodNameT, paramsA[2:]...)

		return rvr

	case "new":
		if len(paramsA) < 1 {
			return fmt.Errorf("参数不够")
		}

		vs1 := tk.ToStr(paramsA[0])

		p := objA.(*xie.XieVM)

		switch vs1 {
		case "window", "webView2":
			return newWindowWebView2(p, paramsA[1:])
		}

		return fmt.Errorf("不支持的创建类型：%v", vs1)

	case "close":
		if len(paramsA) < 1 {
			return fmt.Errorf("参数不够")
		}

		switch nv := paramsA[0].(type) {
		case zenity.ProgressDialog:
			nv.Close()
		}

		return ""

	case "showInfo":
		if len(paramsA) < 2 {
			return fmt.Errorf("参数不够")
		}
		return showInfoGUI(tk.ToStr(paramsA[0]), tk.ToStr(paramsA[1]), paramsA[2:]...)

	case "showError":
		if len(paramsA) < 2 {
			return fmt.Errorf("参数不够")
		}
		return showErrorGUI(tk.ToStr(paramsA[0]), tk.ToStr(paramsA[1]), paramsA[2:]...)

	case "getConfirm":
		if len(paramsA) < 2 {
			return fmt.Errorf("参数不够")
		}
		return getConfirmGUI(tk.ToStr(paramsA[0]), tk.ToStr(paramsA[1]), paramsA[2:]...)
	case "getInput":
		// if len(paramsA) < 2 {
		// 	return fmt.Errorf("参数不够")
		// }
		return getInputGUI(tk.InterfaceToStringArray(paramsA)...)
	case "selectFile":
		// if len(paramsA) < 2 {
		// 	return fmt.Errorf("参数不够")
		// }
		return selectFileGUI(tk.InterfaceToStringArray(paramsA)...)
	case "selectFileToSave":
		// if len(paramsA) < 2 {
		// 	return fmt.Errorf("参数不够")
		// }
		return selectFileToSaveGUI(tk.InterfaceToStringArray(paramsA)...)
	case "getActiveDisplayCount":
		return screenshot.NumActiveDisplays()
	case "getScreenResolution":
		var paraArgsT []string = []string{}

		for i := 0; i < len(paramsA); i++ {
			paraArgsT = append(paraArgsT, tk.ToStr(paramsA[i]))
		}

		pT := objA.(*xie.XieVM)

		formatT := pT.GetSwitchVarValue(pT.Running, paraArgsT, "-format=", "")

		idxStrT := pT.GetSwitchVarValue(pT.Running, paraArgsT, "-index=", "0")

		idxT := tk.StrToInt(idxStrT, 0)

		rectT := screenshot.GetDisplayBounds(idxT)

		if formatT == "" {
			return []interface{}{rectT.Max.X, rectT.Max.Y}
		} else if formatT == "raw" || formatT == "rect" {
			return rectT
		} else if formatT == "json" {
			return tk.ToJSONX(rectT, "-sort")
		}

		return []interface{}{rectT.Max.X, rectT.Max.Y}
	case "showProcess":
		var paraArgsT []string = []string{}

		for i := 0; i < len(paramsA); i++ {
			paraArgsT = append(paraArgsT, tk.ToStr(paramsA[i]))
		}

		optionsT := []zenity.Option{}

		titleT := tk.GetSwitch(paraArgsT, "-title=", "")

		if titleT != "" {
			optionsT = append(optionsT, zenity.Title(titleT))
		}

		okButtonT := tk.GetSwitch(paraArgsT, "-ok=", "")

		if titleT != "" {
			optionsT = append(optionsT, zenity.OKLabel(okButtonT))
		}

		cancelButtonT := tk.GetSwitch(paraArgsT, "-cancel=", "")

		if titleT != "" {
			optionsT = append(optionsT, zenity.CancelLabel(cancelButtonT))
		}

		if tk.IfSwitchExistsWhole(paraArgsT, "-noCancel") {
			optionsT = append(optionsT, zenity.NoCancel())
		}

		if tk.IfSwitchExistsWhole(paraArgsT, "-modal") {
			optionsT = append(optionsT, zenity.Modal())
		}

		if tk.IfSwitchExistsWhole(paraArgsT, "-pulsate") {
			optionsT = append(optionsT, zenity.Pulsate())
		}

		maxT := tk.GetSwitch(paraArgsT, "-max=", "")

		if maxT != "" {
			optionsT = append(optionsT, zenity.MaxValue(tk.ToInt(maxT, 100)))
		}

		dlg, errT := zenity.Progress(optionsT...)
		if errT != nil {
			return fmt.Errorf("创建进度框失败（failed to create progress dialog）：%v", errT)
		}

		return dlg

	case "newWindow":
		return newWindowWebView2(objA, paramsA)

	default:
		return fmt.Errorf("未知方法")
	}

	return ""
}

func initGUI() error {
	// applicationPathT := tk.GetApplicationPath()

	// osT := tk.GetOSName()

	// if tk.Contains(osT, "inux") {
	// } else if tk.Contains(osT, "arwin") {
	// } else {
	// 	// _, errT := exec.LookPath("sciter.dll")

	// 	// // tk.Pln("LookPath", errT)

	// 	// if errors.Is(errT, exec.ErrDot) {
	// 	// 	errT = nil
	// 	// }

	// 	// if errT != nil {
	// 	// 	if tk.IfFileExists("sciter.dll") || tk.IfFileExists(filepath.Join(applicationPathT, "sciter.dll")) {

	// 	// 	} else {
	// 	// 		tk.Pl("初始化WEB图形界面环境……")
	// 	// 		rs := tk.DownloadFile("http://xie.topget.org/pub/sciter.dll", applicationPathT, "sciter.dll")

	// 	// 		if tk.IsErrorString(rs) {
	// 	// 			return fmt.Errorf("初始化图形界面编程环境失败")
	// 	// 		}
	// 	// 	}
	// 	// }
	// }

	// dialog.Do_init()
	// window.Do_init()

	return nil
}

func showInfoGUI(titleA string, formatA string, messageA ...interface{}) interface{} {
	rs, errT := dlgs.Info(titleA, fmt.Sprintf(formatA, messageA...))

	if errT != nil {
		return errT
	}

	return rs
}

func getConfirmGUI(titleA string, formatA string, messageA ...interface{}) interface{} {
	flagT, errT := dlgs.Question(titleA, fmt.Sprintf(formatA, messageA...), true)
	if errT != nil {
		return errT
	}

	return flagT
}

func showErrorGUI(titleA string, formatA string, messageA ...interface{}) interface{} {
	rs, errT := dlgs.Error(titleA, fmt.Sprintf(formatA, messageA...))
	if errT != nil {
		return errT
	}

	return rs
}

// mt $pln $guiG selectFileToSave -confirmOverwrite -title=保存文件…… -default=c:\test\test.txt `-filter=[{"Name":"Go and TextFiles", "Patterns":["*.go","*.txt"], "CaseFold":true}]`

func selectFileToSaveGUI(argsA ...string) interface{} {
	optionsT := []zenity.Option{}

	optionsT = append(optionsT, zenity.ShowHidden())

	titleT := tk.GetSwitch(argsA, "-title=", "")

	if titleT != "" {
		optionsT = append(optionsT, zenity.Title(titleT))
	}

	defaultT := tk.GetSwitch(argsA, "-default=", "")

	if defaultT != "" {
		optionsT = append(optionsT, zenity.Filename(defaultT))
	}

	if tk.IfSwitchExistsWhole(argsA, "-confirmOverwrite") {
		optionsT = append(optionsT, zenity.ConfirmOverwrite())
	}

	filterStrT := tk.GetSwitch(argsA, "-filter=", "")

	// tk.Plv(filterStrT)

	var filtersT zenity.FileFilters

	if filterStrT != "" {

		errT := jsoniter.Unmarshal([]byte(filterStrT), &filtersT)

		if errT != nil {
			return errT
		}

		optionsT = append(optionsT, filtersT)
	}

	rs, errT := zenity.SelectFileSave(optionsT...)

	if errT != nil {
		if errT == zenity.ErrCanceled {
			return nil
		}

		return errT
	}

	return rs
}

func selectFileGUI(argsA ...string) interface{} {
	optionsT := []zenity.Option{}

	optionsT = append(optionsT, zenity.ShowHidden())

	titleT := tk.GetSwitch(argsA, "-title=", "")

	if titleT != "" {
		optionsT = append(optionsT, zenity.Title(titleT))
	}

	defaultT := tk.GetSwitch(argsA, "-default=", "")

	if defaultT != "" {
		optionsT = append(optionsT, zenity.Filename(defaultT))
	}

	filterStrT := tk.GetSwitch(argsA, "-filter=", "")

	var filtersT zenity.FileFilters

	if filterStrT != "" {

		errT := jsoniter.Unmarshal([]byte(filterStrT), &filtersT)

		if errT != nil {
			return errT
		}

		optionsT = append(optionsT, filtersT)
	}

	rs, errT := zenity.SelectFile(optionsT...)

	if errT != nil {
		if errT == zenity.ErrCanceled {
			return nil
		}

		return errT
	}

	return rs
}

func getInputGUI(argsA ...string) interface{} {
	optionsT := []zenity.Option{}

	optionsT = append(optionsT, zenity.ShowHidden())

	titleT := tk.GetSwitch(argsA, "-title=", "")

	if titleT != "" {
		optionsT = append(optionsT, zenity.Title(titleT))
	}

	defaultT := tk.GetSwitch(argsA, "-default=", "")

	if defaultT != "" {
		optionsT = append(optionsT, zenity.EntryText(defaultT))
	}

	hideTextT := tk.IfSwitchExistsWhole(argsA, "-hideText")
	if hideTextT {
		optionsT = append(optionsT, zenity.HideText())
	}

	modalT := tk.IfSwitchExistsWhole(argsA, "-modal")
	if modalT {
		optionsT = append(optionsT, zenity.Modal())
	}

	textT := tk.GetSwitch(argsA, "-text=", "")

	okLabelT := tk.GetSwitch(argsA, "-okLabel=", "")

	if okLabelT != "" {
		optionsT = append(optionsT, zenity.OKLabel(okLabelT))
	}

	cancelLabelT := tk.GetSwitch(argsA, "-cancelLabel=", "")

	if cancelLabelT != "" {
		optionsT = append(optionsT, zenity.CancelLabel(cancelLabelT))
	}

	extraButtonT := tk.GetSwitch(argsA, "-extraButton=", "")

	if extraButtonT != "" {
		optionsT = append(optionsT, zenity.ExtraButton(extraButtonT))
	}

	rs, errT := zenity.Entry(textT, optionsT...)

	if errT != nil {
		if errT == zenity.ErrCanceled {
			return nil
		}

		if errT == zenity.ErrExtraButton {
			return fmt.Errorf("extraButton")
		}

		return errT
	}

	return rs
}
