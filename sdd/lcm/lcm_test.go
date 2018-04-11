package lcm_calc

import (
	"testing"
)

func TestLcmCalc(t *testing.T) {

	tests := []struct {
		name  string
		args  []int
		wants int
	}{
		{
			name:  "Two numbers",
			args:  []int{3, 5},
			wants: 15,
		},
		{
			name:  "Three numbers",
			args:  []int{3, 5, 9},
			wants: 45,
		},
		{
			name:  "Large numbers",
			args:  []int{38756, 5487, 93294},
			wants: 3306559720428,
		},
		{
			name:  "Larger array",
			args:  []int{5, 17, 90, 23, 23, 43, 86, 75, 109, 203},
			wants: 167409562950,
		},
		{
			name:  "Numbers with zero",
			args:  []int{45, 0, 5, 9},
			wants: 0,
		},
		{
			name:  "Negative Numbers",
			args:  []int{-8, 2, -5, 4},
			wants: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if lcm := LcmCalc(tt.args); lcm != tt.wants {
				t.Errorf("LcmCalc(%v) = %v, wants %v", tt.args, lcm, tt.wants)
			}
		})
	}
}
