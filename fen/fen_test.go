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
	f.Add("4kb1r/p2rqppp/5n2B2p1B1/4P3/1Q6/PPP2PPP/2K4R w - - 0 1")
	f.Add("3k4/2p2p2/1p5p/p1p1P1p1/P1Pn2P1/1P3P1P/1B3K2/8 w - - 0 30")
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
