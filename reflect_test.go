package gotils

import (
	"reflect"
	"testing"
)

func Test_FromPtr(t *testing.T) {
	t1 := "qwe"
	t2 := &t1
	t3 := &t2
	t4 := &t3
	t5 := &t4

	got, err := FromPtr(t5)
	if err != nil {
		t.Errorf("fromPtr() error = %v, wantErr false", err)
		return
	}
	if !reflect.DeepEqual(got, t1) {
		t.Errorf("fromPtr() got = %v, want %v", got, t1)
	}

	got, err = FromPtr(t1)
	if err != nil {
		t.Errorf("fromPtr() error = %v, wantErr false", err)
		return
	}

	if !reflect.DeepEqual(got, t1) {
		t.Errorf("fromPtr() got = %v, want %v", got, t1)
		return
	}

	t5 = nil
	t7 := &t5
	got, err = FromPtr(t7)
	if err == nil {
		t.Errorf("fromPtr() error = %v, wantErr false", err)
		return
	}
}
