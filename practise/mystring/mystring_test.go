package string

import "testing"

func Test(t *testing.T) {
	var tests = []struct{
		s, want string 
	}{
		{"1234", "4321"},
		{"12345", "54321"},
		{"好好学习", "习学好好"},
		{"", ""},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}