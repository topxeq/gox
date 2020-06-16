// +build !windows

package main

import (
	"github.com/topxeq/qlang"
)

func GetSystemMetrics(nIndex int) int {
	return 0
}

func GetScreenResolution() (int, int) {
	return 0, 0
}

func InitSysspec() {
	var sysspecExports = map[string]interface{}{
		"GetScreenResolution": GetScreenResolution,
		"GetSystemMetrics":    GetSystemMetrics,
	}

	qlang.Import("sysspec", sysspecExports)
}
