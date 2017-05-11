package main

import (
	"github.com/imnotanderson/X/agent"
	"github.com/imnotanderson/X/launcher"
)

func main() {
	l := &launcher.Launcher{}
	l.Start(
		agent.Module,
	)

}
