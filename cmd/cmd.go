package cmd

import (
	"bufio"
	"github.com/imnotanderson/X/log"
	"io"
	"os"
	"strings"
)

type Cmd struct {
	in   io.Reader
	fmap map[string]func([]string)
}

var Module *Cmd = &Cmd{
	in:   os.Stdin,
	fmap: map[string]func([]string){},
}

func (c *Cmd) Init() {

}

func (c *Cmd) Run(closeSign <-chan struct{}) {
	go c.handleInput()
}

func (c *Cmd) handleInput() {
	for {
		reader := bufio.NewReader(c.in)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line[:len(line)-1], "\r")
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		f := c.fmap[args[0]]
		if f == nil {
			log.Errorf("no func %v", args[0])
		} else {
			f(args[1:])
		}
	}
}

func (c *Cmd) RegFunc(name string, f func([]string)) {
	if c.fmap[name] != nil {
		log.Errorf("same func %v", name)
		return
	}
	c.fmap[name] = f
}
