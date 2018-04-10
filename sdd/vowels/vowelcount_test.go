package vowels

import (
	"reflect"
	"testing"
)

func TestVowelCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Single sentence",
			args: args{s: "This is just a simple test sentence."},
			want: map[string]int{"a": 1, "e": 5, "i": 3, "o": 0, "u": 1},
		},
		{
			name: "Two sentences with all the vowels",
			args: args{s: "These two sentences have all the vowels. What do you think of that?"},
			want: map[string]int{"a": 4, "e": 8, "i": 1, "o": 5, "u": 1},
		},
		{
			name: "Sentence with a special character",
			args: args{s: "This sentence has a special 漢 character! Whoa!"},
			want: map[string]int{"a": 6, "e": 5, "i": 2, "o": 1, "u": 0},
		},
		{
			name: "No vowels",
			args: args{s: "xcvbnqwrtpsdfgh"},
			want: map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0},
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VowelCount(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
