package fibonaccisequence

import (
	"reflect"
	"testing"
)

func benchmarkFibonacciSequence(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciSequence(i)
	}
}
func BenchmarkFibonacciSequence1(b *testing.B)  { benchmarkFibonacciSequence(1, b) }
func BenchmarkFibonacciSequence10(b *testing.B) { benchmarkFibonacciSequence(10, b) }
func BenchmarkFibonacciSequence15(b *testing.B) { benchmarkFibonacciSequence(15, b) }
func BenchmarkFibonacciSequence20(b *testing.B) { benchmarkFibonacciSequence(20, b) }

func BenchmarkFibonacciSequence25(b *testing.B) { benchmarkFibonacciSequence(25, b) }

func BenchmarkFibonacciSequence30(b *testing.B) { benchmarkFibonacciSequence(30, b) }
func BenchmarkFibonacciSequence40(b *testing.B) { benchmarkFibonacciSequence(40, b) }

func Test_FibonacciSequence(t *testing.T) {
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
			if got := FibonacciSequence(tt.args); !reflect.DeepEqual(got, tt.wants) {
				t.Errorf("FibonacciSequence(%d) = %v, wants %v", tt.args, got, tt.wants)
			}
		})
	}
}
