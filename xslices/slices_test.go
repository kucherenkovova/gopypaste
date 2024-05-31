package xslices_test

import (
	"reflect"
	"testing"

	"github.com/kucherenkovova/gopypaste/xslices"
)

func TestUnique(t *testing.T) {
	type testCase struct {
		name string
		in   []string
		want []string
	}

	tests := []testCase{
		{
			name: "empty list",
			in:   []string{},
			want: []string{},
		},
		{
			name: "nil",
			in:   nil,
			want: nil,
		},
		{
			name: "no duplicates",
			in:   []string{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "consequent duplicates",
			in:   []string{"a", "a", "b", "b", "c", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "non-consequent duplicates",
			in:   []string{"a", "b", "a", "c", "b", "c"},
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xslices.Unique(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
