package objectify

import (
	"reflect"
	"testing"
)

func TestObjectify(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []*Object
		err   bool
	}{
		{
			name: "Simple",
			input: []string{
				"Orange, fruit, 7.5, orange",
				"Apple, fruit, 6.5, red",
			},
			want: []*Object{
				&Object{"Orange", "fruit", 7.5, "orange"},
				&Object{"Apple", "fruit", 6.5, "red"},
			},
			err: false,
		},
		{
			name:  "Empty",
			input: []string{},
			want:  []*Object{},
			err:   false,
		},
		{
			name:  "Bad string",
			input: []string{"bad input"},
			want:  nil,
			err:   true,
		},
		{
			name:  "Bad size type",
			input: []string{"Obj, sometime, badSize, somecolor"},
			want:  nil,
			err:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Objectify(tt.input)
			if tt.err && err == nil {
				t.Errorf("Expected error on input: %v", tt.input)
			}
			if !tt.err && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Objectify(s []string) = %v, want %v", got, tt.want)
			}
		})
	}
}
