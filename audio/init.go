package audio

import "golang.org/x/sys/windows"

var audio *windows.LazyDLL

// Init -> audio
func Init(path string) {
	audio = windows.NewLazyDLL(path + "csfml-audio-2.dll")
}