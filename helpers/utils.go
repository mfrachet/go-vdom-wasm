package vnh

import "reflect"

func IsNil(toVerify interface{}) bool {
	return reflect.ValueOf(toVerify).IsNil()
}

func NotNil(toVerify interface{}) bool {
	return !IsNil(toVerify)
}
