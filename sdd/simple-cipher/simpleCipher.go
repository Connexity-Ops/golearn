package simplecipher

// simpleEncrypt takes a string like "Hello World" and returns an array of integers that are the
// numeric equivalent of the letters. Example A=1, B=2, C=3, ...
// If the character is a number then the return code will be 3 + the number,
// for example "2" will return 32.
// If the character is a space then the return code is a 0.
// All punctuation and special characters are ignored.

func simpleEncrpyt(r string) []int {
	// TODO return an array of integer(s) mapping to string.
	var ret []int
	for i := 0; i < len(r); i++ {
		c := int(r[i])

		switch {
		case c > 96 && c < 123:
			ret = append(ret, c-96)
		case c > 64 && c < 91:
			ret = append(ret, c-64)
		case c > 47 && c < 58:
			ret = append(ret, c-47+29)
		case c == 32:
			ret = append(ret, c-32)
		}
	}
	return (ret)
}
