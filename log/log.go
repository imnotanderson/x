package log

import l "log"

type log struct {
}

var log *log = &log{}

func Error(agrs ...interface{}) {
	l.Print("err:", agrs...)
}

func Errorf(format string, args ...interface{}) {
	l.Printf("err:"+format, args...)
}

func Infof(format string, args ...interface{}) {
	l.Printf("info:"+format, args...)
}
