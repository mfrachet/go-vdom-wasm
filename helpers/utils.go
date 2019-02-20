package vnh

import (
	"reflect"
)

func IsNil(toVerify interface{}) bool {
	if toVerify == nil {
		return true
	}

	return reflect.ValueOf(toVerify).IsNil()
}

func NotNil(toVerify interface{}) bool {
	return !IsNil(toVerify)
}
