package log

import (
	"fmt"
	"github.com/fdev-ci/golang-plugin-sdk/util/ansi"
	"time"
)

func Info(v ...interface{}) {
	fmt.Printf("%v%v [info]  %v %v \n", ansi.INFO, time.Now().Format("2006-01-02 15:04:05"), v, ansi.RESET)
}

func Warn(v ...interface{}) {
	fmt.Printf("%v%v [warn] %v %v \n", ansi.WARN, time.Now().Format("2006-01-02 15:04:05"), v, ansi.RESET)
}

func Error(v ...interface{}) {
	fmt.Printf("%v%v [error] %v %v \n", ansi.ERROR, time.Now().Format("2006-01-02 15:04:05"), v, ansi.RESET)
}

func Debug(v ...interface{}) {
	fmt.Printf("%v%v [debug] %v %v \n", ansi.DEBUG, time.Now().Format("2006-01-02 15:04:05"), v, ansi.RESET)
}
