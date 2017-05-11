package log

import "fmt"

func Errorf(format string, args ...interface{}) {
	fmt.Printf("err:"+format+"\n", args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf("info:"+format+"\n", args...)
}
