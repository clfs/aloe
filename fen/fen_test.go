package fen

import (
	"testing"

	"github.com/clfs/aloe/chess"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		pos  chess.Position
		want string
	}{
		{
			pos:  chess.NewPosition(),
			want: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
	}
	for _, c := range cases {
		got, err := Encode(c.pos)
		if err != nil {
			t.Errorf("unable to decode %q: %v", c.want, err)
		}
		if c.want != got {
			t.Errorf("want %q, got %q", c.want, got)
		}
	}
}

func FuzzRoundTrip(f *testing.F) {
	f.Add(StartingPosition)
	f.Fuzz(func(t *testing.T, old string) {
		pos, err := Decode(old)
		if err != nil {
			t.Skip() // Invalid input.
		}

		new, err := Encode(pos)
		if err != nil {
			t.Errorf("encode failed after decoding: %v", err)
		}

		if old != new {
			t.Errorf("changed after round trip: old %q, new %q", old, new)
		}
	})
}
