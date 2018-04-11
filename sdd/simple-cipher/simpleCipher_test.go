package simpleCipher

import (
	"reflect"
	"testing"
)

func Test_simpleEncrpyt(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name  string
		args  string
		wants []int
	}{
		{
			name:  "Simple String",
			args:  "Hello World",
			wants: []int{8, 5, 12, 12, 15, 0, 23, 15, 18, 12, 4},
		},
		{
			name:  "String with numbers",
			args:  "Meeting time is 3:15pm",
			wants: []int{13, 5, 5, 20, 9, 14, 7, 0, 20, 9, 13, 5, 0, 9, 19, 0, 33, 31, 35, 16, 13},
		},
		{
			name:  "Special characters only",
			args:  ".!@#%$^&",
			wants: []int{},
		},
		{
			name:  "Longer sentence",
			args:  "It's raining cats and dogs! Go get an umbrella!",
			wants: []int{9, 20, 19, 0, 18, 1, 9, 14, 9, 14, 7, 0, 3, 1, 20, 19, 0, 1, 14, 4, 0, 4, 15, 7, 19, 0, 7, 15, 0, 7, 5, 20, 0, 1, 14, 0, 21, 13, 2, 18, 5, 12, 12, 1},
		},
		{
			name:  "Nothing",
			args:  "",
			wants: []int{},
		},
		{
			name:  "All spaces",
			args:  "      ",
			wants: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name:  "Numbers Only",
			args:  "0123456789",
			wants: []int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleEncrpyt(tt.args); !reflect.DeepEqual(got, tt.wants) {
				t.Errorf("toNum() = %v, wants %v", got, tt.wants)
			}
		})
	}
}
