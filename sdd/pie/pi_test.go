package pi

import (
	"testing"
)

func TestPi(t *testing.T) {
	var piTests = []struct {
		input    int
		expected int
	}{
		{1, 3},
		{2, 1},
		{3, 4},
		{200, 9},
		{99999, 6},
	}

	for _, tt := range piTests {
		actual := Pi(tt.input)
		if actual != tt.expected {
			t.Errorf("Pi(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}
}
