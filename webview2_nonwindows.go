//go:build !windows
// +build !windows

package gox

func newWebView2(optsA ...string) interface{} {
	return nil
}
