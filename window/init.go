package window

import (
	"runtime"

	"golang.org/x/sys/windows"
)


var window *windows.LazyDLL

// Init -> window
func Init(path string) {
	runtime.LockOSThread()
	window = windows.NewLazyDLL(path + "csfml-window-2.dll")
}