package log

import (
	"fmt"
)

func Errorf(format string, args ...interface{}) {
	fmt.Printf("err:"+format+"\n", args...)
	//debug.PrintStack()
}

func Infof(format string, args ...interface{}) {
	fmt.Printf("info:"+format+"\n", args...)
	//debug.PrintStack()
}

func Debugf(format string, args ...interface{}) {
	//debug.PrintStack()
	fmt.Printf("debug:"+format+"\n", args...)
}
