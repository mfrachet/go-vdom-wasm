package vn

type KeyIdentifier interface {
	SetValue(string)
	GetValue() string
}

type Key struct {
	Value string
}

func (Key *Key) SetValue(value string) {
	Key.Value = value
}

func (key *Key) GetValue() string {
	return key.Value
}
