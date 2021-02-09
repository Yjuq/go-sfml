package sfml

import (
	"github.com/yjuq/go-sfml/audio"
	"github.com/yjuq/go-sfml/graphics"
	"github.com/yjuq/go-sfml/network"
	"github.com/yjuq/go-sfml/system"
	"github.com/yjuq/go-sfml/window"
)

// Init -> sfml
func Init(path string) {
	audio.Init(path)
	graphics.Init(path)
	network.Init(path)
	system.Init(path)
	window.Init(path)
}