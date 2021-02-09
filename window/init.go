package window

import "golang.org/x/sys/windows"

var window *windows.LazyDLL

// Init -> window
func Init(path string) {
	window = windows.NewLazyDLL(path + "csfml-window-2.dll")
}