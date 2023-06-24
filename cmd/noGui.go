//go:build linux
// +build linux

package main

import (
	"fmt"
)

func guiHandler(actionA string, objA interface{}, dataA interface{}, paramsA ...interface{}) interface{} {
	return fmt.Errorf("GUI engined disabled")
}
