package launcher

import (
	"testing"
	"time"
)

type mod struct {
	name string
}

func (m *mod) Run(closeSign <-chan struct{}) {
	println(m.name, " run")
	<-closeSign
	println(m.name, " close")
}

func (m *mod) Init() {
	println(m.name, " init")
}

func TestRun(t *testing.T) {
	x := &Launcher{}
	x.Start(
		&mod{"mod1"},
		&mod{"mod2"},
		&mod{"mod3"},
		&mod{"mod4"},
		&mod{"mod5"},
		&mod{"mod6"},
	)
	<-time.After(time.Second * 3)
	x.Close()
	println("close")
}
