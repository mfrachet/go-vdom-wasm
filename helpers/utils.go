package vnh

func IsNil(toVerify interface{}) bool {
	return toVerify == nil
}

func NotNil(toVerify interface{}) bool {
	return !IsNil(toVerify)
}
