package vn

type KeyIdentifier interface {
	SetValue(string)
	GetValue() string
}

type Key struct {
	value string
}

func (Key *Key) SetValue(value string) {
	Key.value = value
}

func (key *Key) GetValue() string {
	return key.value
}
