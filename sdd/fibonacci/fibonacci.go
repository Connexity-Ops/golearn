package fibonaccisequence

// fibonacciSequence takes an integer e.g. "1, 2, 5, 10, 30" and return a sequence of integer(s)
// of the famous Fibonacci Sequence
// Return an "EMPTY" slice for negative numbers

// FibonacciSequence assembly the sequence slice
func FibonacciSequence(n int) []int {
	ret := []int{}
	for i := 0; i <= n; i++ {
		ret = append(ret, fibonacci(i))
	}
	return ret
}

// fibonacci return the sequence elements
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
