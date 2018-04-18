package lyrics

import "testing"

func TestGetLyrics(t *testing.T) {
	type args struct {
		b string
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check a short song",
			args: args{b: "The Champs", s: "Tequila"},
			want: "Tequila\n\nTequila\n\nTequila",
		},
		{
			name: "Handle an API error",
			args: args{b: "Mike Jones", s: "I totally wrote a song! Not."},
			want: "No lyrics found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLyrics(tt.args.b, tt.args.s); got != tt.want {
				t.Errorf("GetLyrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
