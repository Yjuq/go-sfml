package graphics

import "golang.org/x/sys/windows"

var graphics *windows.LazyDLL

func Init(path string) {
	graphics = windows.NewLazyDLL(path + "csfml-graphics-2.dll")
}