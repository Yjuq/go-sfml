package sfml

import (
	"github.com/Yjuq/go-sfml/audio"
	"github.com/Yjuq/go-sfml/graphics"
	"github.com/Yjuq/go-sfml/network"
	"github.com/Yjuq/go-sfml/system"
	"github.com/Yjuq/go-sfml/window"
)

func Init(path string) {
	audio.Init(path)
	graphics.Init(path)
	network.Init(path)
	system.Init(path)
	window.Init(path)
}