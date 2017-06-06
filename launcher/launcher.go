package launcher

import (
	"github.com/imnotanderson/X/utils"
	"sync"
)

type IModule interface {
	Init()
	Run(closeSign <-chan struct{})
}

type Launcher struct {
	moduleSigns []chan struct{}
	wg          *sync.WaitGroup
}

func (l *Launcher) Start(modules ...IModule) {
	l.wg = &sync.WaitGroup{}
	l.moduleSigns = make([]chan struct{}, len(modules))
	for _, m := range modules {
		m.Init()
	}
	for i, _ := range modules {
		m := modules[i]
		c := make(chan struct{})
		l.moduleSigns[i] = c
		l.wg.Add(1)
		go func() {
			defer utils.PrintPanicStack()
			defer l.wg.Done()
			m.Run(c)
		}()
	}
	l.wg.Wait()
}

func (l *Launcher) Close() {
	for _, v := range l.moduleSigns {
		close(v)
	}
	l.wg.Wait()
}
