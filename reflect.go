package gotils

import (
	"fmt"
	"reflect"
)

var ErrNilPointer = fmt.Errorf("nil pointer")

func FromPtr[T any](v any) (t T, err error) {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Ptr {
		if vValue.IsNil() {
			return t, ErrNilPointer
		}
		vValue = vValue.Elem()
		t, err = FromPtr[T](vValue.Interface())
		return t, err
	}

	return v.(T), nil
}
