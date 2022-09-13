package chess

import "testing"

func TestSquare_IsAdjacentTo(t *testing.T) {
	cases := []struct {
		s1, s2 Square
		want   bool
	}{
		{A1, A1, false},
		{A1, A2, true},
		{A1, A3, false},
		{A1, B1, true},
		{A1, B2, true},
		{E4, D5, true},
	}
	for _, tc := range cases {
		got := tc.s1.IsAdjacentTo(tc.s2)
		if tc.want != got {
			t.Errorf("%v, %v: want %v, got %v", tc.s1, tc.s2, tc.want, got)
		}
	}
}

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
