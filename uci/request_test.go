package uci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRequestPosition_UnmarshalText(t *testing.T) {
	cases := []struct {
		text []byte
		want RequestPosition
	}{
		{
			text: []byte("position startpos"),
			want: RequestPosition{
				FEN:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
				Moves: nil,
			},
		},
		{
			text: []byte("position startpos moves e2e4 e7e5"),
			want: RequestPosition{
				FEN:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
				Moves: []string{"e2e4", "e7e5"},
			},
		},
		{
			text: []byte("position rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"),
			want: RequestPosition{
				FEN:   "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
				Moves: nil,
			},
		},
		{
			text: []byte("position rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2 moves e2e4 e7e5"),
			want: RequestPosition{
				FEN:   "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
				Moves: []string{"e2e4", "e7e5"},
			},
		},
	}
	for i, c := range cases {
		var got RequestPosition
		if err := got.UnmarshalText(c.text); err != nil {
			t.Errorf("case %d: %v", i, err)
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("%d: (-want +got)\n%s", i, diff)
		}
	}
}
