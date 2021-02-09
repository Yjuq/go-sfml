package network

import "golang.org/x/sys/windows"

var network *windows.LazyDLL

// Init -> network
func Init(path string) {
	network = windows.NewLazyDLL(path + "csfml-network-2.dll")
}