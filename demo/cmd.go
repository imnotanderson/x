package main

import (
	"github.com/imnotanderson/X/cmd"
	"github.com/imnotanderson/X/launcher"
)

func main() {
	(&launcher.Launcher{}).Start(cmd.Module)
	select {}
}
