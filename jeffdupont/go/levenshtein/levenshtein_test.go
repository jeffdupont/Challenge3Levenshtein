package levenshtein

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "GUMBO : GAMBOL", args: args{"GUMBO", "GAMBOL"}, want: 2},
		{name: "KITTEN : SITTING", args: args{"KITTEN", "SITTING"}, want: 3},
		{name: "SATURDAY : SUNDAY", args: args{"SATURDAY", "SUNDAY"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
