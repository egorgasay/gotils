package gotils

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type E any
	type els[S interface{ ~[]E }] struct {
		arr S
	}
	type testCase[S interface{ ~[]E }] struct {
		name string
		args els[S]
		want els[S]
	}
	tests := []testCase[[]E]{
		{
			name: "Reverse []int",
			args: els[[]E]{
				arr: []E{1, 2, 3, 4, 5, 6},
			},
			want: els[[]E]{
				arr: []E{6, 5, 4, 3, 2, 1},
			},
		},
		{
			name: "Reverse []any",
			args: els[[]E]{
				arr: []E{"1", 2, "qdwdq"},
			},
			want: els[[]E]{
				arr: []E{"qdwdq", 2, "1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.arr)

			if !reflect.DeepEqual(tt.args.arr, tt.want.arr) {
				t.Fail()
			}
		})
	}
}
