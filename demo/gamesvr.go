package main

import (
	"github.com/imnotanderson/X/game"
	"github.com/imnotanderson/X/launcher"
)

func main() {
	l := launcher.Launcher{}
	l.Start(
		game.Module,
	)
}
