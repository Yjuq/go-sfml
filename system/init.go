package system

import "golang.org/x/sys/windows"

var system *windows.LazyDLL

// Init -> system
func Init(path string) {
	system = windows.NewLazyDLL(path + "csfml-system-2.dll")
	clock()
}