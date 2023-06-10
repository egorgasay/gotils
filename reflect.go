package gotils

import (
	"fmt"
	"reflect"
)

var ErrNilPointer = fmt.Errorf("nil pointer")

func FromPtr(v any) (any, error) {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Ptr {
		if vValue.IsNil() {
			return nil, ErrNilPointer
		}
		vValue = vValue.Elem()
		return FromPtr(vValue.Interface())
	}

	return v, nil
}
