package uci

import (
	"testing"

	"github.com/clfs/aloe/fen"
)

func TestParsePosition(t *testing.T) {
	cases := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{
			in:   "position startpos",
			want: fen.StartingFEN,
		},
		{
			in:   "position startpos moves",
			want: fen.StartingFEN,
		},
		{
			in:   "position startpos moves e2e4",
			want: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
			want: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1 moves e2e4",
			want: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1 moves e2e4 e7e5",
			want: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			in:      "position",
			wantErr: true,
		},
		{
			in:      "position startpos moves e5e6",
			wantErr: true,
		},
		{
			in:      "position fen",
			wantErr: true,
		},
		{
			in:      "position fen rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1 moves a1h8",
			wantErr: true,
		},
	}
	for _, tc := range cases {
		got, err := ParsePosition(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("%q: incorrect error value: %v", tc.in, err)
		}
		if tc.want != got {
			t.Errorf("%q: want %q, got %q", tc.in, tc.want, got)
		}
	}
}
