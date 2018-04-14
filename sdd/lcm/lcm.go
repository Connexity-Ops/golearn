package lcm_calc

import (
	"math"
)

// gcd finds the greatest common denominator of two numbers.
func gcd(x, y int64) int64 {

	for y != 0 {
		x, y = y, x%y
	}

	return x
}

// LcmCalc finds the Least Common Multiple using the gcd function on an array of numbers.
func LcmCalc(numbers []int64) int64 {
	x := numbers[0]
	for _, y := range numbers {
		x = x * y / gcd(x, y)
	}

	return int64(math.Abs(float64(x)))

}
