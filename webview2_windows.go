//go:build windows
// +build windows

package main

import (
	"github.com/jchv/go-webview2"
	"github.com/topxeq/tk"
)

func newWebView2(optsA ...string) interface{} {
	titleT := tk.GetSwitch(optsA, "-title=", "Gox "+versionG)
	widthT := uint(tk.ToInt(tk.GetSwitch(optsA, "-width=", "800"), 800))
	heightT := uint(tk.ToInt(tk.GetSwitch(optsA, "-height=", "600"), 600))
	iconT := uint(tk.ToInt(tk.GetSwitch(optsA, "-icon=", "2"), 2))
	centerT := !tk.IfSwitchExistsWhole(optsA, "-centerFalse")

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  titleT,
			Width:  widthT,
			Height: heightT,
			IconId: iconT, // icon resource id
			Center: centerT,
		},
	})

	return w
}
