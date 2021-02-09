package graphics

import "golang.org/x/sys/windows"

var graphics *windows.LazyDLL

// Init -> graphics
func Init(path string) {
	graphics = windows.NewLazyDLL(path + "csfml-graphics-2.dll")
}