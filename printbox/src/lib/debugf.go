package lib

import "fmt"

var Debug = false

func Debugf(f string, v ...interface{}) {
	if !Debug {
		return
	}
	msg := fmt.Sprintf(f, v...)
	fmt.Println(msg)
}

func Logf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	fmt.Println(msg)
}
