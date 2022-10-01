package chess

import "testing"

func TestSquare_String(t *testing.T) {
	cases := []struct {
		s    Square
		want string
	}{
		{A1, "A1"},
		{E4, "E4"},
		{100, "Square(100)"},
	}
	for _, tc := range cases {
		got := tc.s.String()
		if tc.want != got {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}
}
