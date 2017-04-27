package utils

import (
	. "github.com/imnotanderson/X/log"
	"runtime"
)

// 产生panic时的调用栈打印
func PrintPanicStack() {
	if x := recover(); x != nil {
		Error(x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}
	}
}

func CheckErr(err error) bool {
	if err != nil {
		println(err)
		return true
	}
	return false
}
