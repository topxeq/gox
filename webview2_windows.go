//go:build windows
// +build windows

package gox

import (
	"github.com/jchv/go-webview2"
	"github.com/topxeq/tk"
)

func newWebView2(optsA ...string) interface{} {
	if ServerModeG {
		return nil
	}

	titleT := tk.GetSwitch(optsA, "-title=", "Gox "+VersionG)
	widthT := uint(tk.ToInt(tk.GetSwitch(optsA, "-width=", "800"), 800))
	heightT := uint(tk.ToInt(tk.GetSwitch(optsA, "-height=", "600"), 600))
	iconT := uint(tk.ToInt(tk.GetSwitch(optsA, "-icon=", "2"), 2))
	centerT := tk.IfSwitchExistsWhole(optsA, "-center")
	debugT := tk.IfSwitchExistsWhole(optsA, "-debug")

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     debugT,
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
