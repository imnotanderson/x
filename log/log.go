package log

import (
	"fmt"
)

func Errorf(format string, args ...interface{}) {
	fmt.Printf("err:"+format+"\n", args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf("info:"+format+"\n", args...)
}

func Debugf(format string, args ...interface{}) {
	//debug.PrintStack()
	fmt.Printf("debug:"+format+"\n", args...)
}
