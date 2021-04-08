package root

import (
	"errors"
	"reflect"
)

type Header struct {
	Header  string
	Content string
}

type Request struct {
	Name    string
	Method  string
	Headers []Header
	Body    string
	URL     string
}

func TestStructType(testInterface interface{}, expectedType string) error {
	currentType := reflect.TypeOf(testInterface)

	if expectedType != currentType.Name() {
		error_message := "Expected " + expectedType + " got " + currentType.Name()
		err := errors.New(error_message)
		return err
	}

	return nil
}
