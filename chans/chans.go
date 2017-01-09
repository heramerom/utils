package chans

import (
	"fmt"
	"reflect"
)

func SafeClose(ch interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("close Chan error! %v", x)
		}
	}()
	reflect.ValueOf(ch).Close()
	return
}

func SafeSend(ch interface{}, value interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("send Chan error! %v", x)
		}
	}()
	reflect.ValueOf(ch).Send(reflect.ValueOf(value))
	return
}
