//go:build !windows
// +build !windows

package main

func newWebView2(optsA ...string) interface{} {
	return nil
}
