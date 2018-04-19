package fibonacci_sequence

import (
	"reflect"
	"testing"
)

func Test_fibonacciSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  int
		wants []int
	}{
		{
			name:  "Short Sequence",
			args:  4,
			wants: []int{0, 1, 1, 2, 3},
		},
		{
			name:  "Larger Sequence",
			args:  10,
			wants: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
		{
			name:  "Zero Sequence",
			args:  0,
			wants: []int{0},
		},
		{
			name:  "Special characters only",
			args:  -15,
			wants: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciSequence(tt.args); !reflect.DeepEqual(got, tt.wants) {
				t.Errorf("fibonacciSequence(%d) = %v, wants %v", tt.args, got, tt.wants)
			}
		})
	}
}
