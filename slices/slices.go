package slices

import "reflect"

func Contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

func ContainsFunc(s interface{}, elem interface{}, equal func(v1 interface{}, v2 interface{}) bool) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if equal(arrV.Index(i).Interface(), elem) {
				return true
			}
		}
	}
	return false
}
