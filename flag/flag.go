package flag

import (
	"reflect"
)

type Flag struct {
	Key string
	// this means is flag will take input or not
	TakesInput bool
	InputType  reflect.Kind

	Data interface{}
	Used bool
}

// Takes pointer
func New(key string, returnValue interface{}) *Flag {
	return &Flag{Key: key, Data: returnValue}
}

func (f *Flag) AddInput(inputType reflect.Kind) *Flag {
	f.TakesInput = true
	f.InputType = inputType
	return f
}
