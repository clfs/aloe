package uci

import (
	"testing"

	"github.com/clfs/aloe/fen"
	"github.com/google/go-cmp/cmp"
)

func TestRequestIsReady_UnmarshalText(t *testing.T) {
	var req RequestIsReady

	text := []byte("isready")
	if err := req.UnmarshalText(text); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	text = []byte("something else")
	if err := req.UnmarshalText(text); err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestRequestGo_UnmarshalText(t *testing.T) {
	cases := []struct {
		text    []byte
		want    RequestGo
		wantErr bool
	}{}

	var req RequestGo

	for _, c := range cases {
		if err := req.UnmarshalText(c.text); c.wantErr != (err != nil) {
			t.Errorf("unexpected error: %v", err)
		}
		if diff := cmp.Diff(c.want, req); diff != "" {
			t.Errorf("%q: (-want, +got)\n%s", c.text, diff)
		}
	}
}

func TestRequestPosition_UnmarshalText(t *testing.T) {
	cases := []struct {
		in   string
		want RequestPosition
	}{
		{
			in:   "position startpos",
			want: RequestPosition{FEN: fen.StartingFEN},
		},
		{
			in:   "position startpos moves e2e4",
			want: RequestPosition{fen.StartingFEN, []string{"e2e4"}},
		},
		{
			in:   "position startpos moves e2e4 e7e5",
			want: RequestPosition{fen.StartingFEN, []string{"e2e4", "e7e5"}},
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			want: RequestPosition{FEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 moves e2e4",
			want: RequestPosition{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", []string{"e2e4"}},
		},
		{
			in:   "position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 moves e2e4 e7e5",
			want: RequestPosition{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", []string{"e2e4", "e7e5"}},
		},
	}

	var req RequestPosition

	for _, c := range cases {
		if err := req.UnmarshalText([]byte(c.in)); err != nil {
			t.Errorf("%q: error: %v", c.in, err)
		}
		if diff := cmp.Diff(c.want, req); diff != "" {
			t.Errorf("%q: (-want, +got)\n%s", c.in, diff)
		}
	}

	invalidCases := []string{
		"position",
		"position fen",
		"position startpos moves",
	}

	for _, c := range invalidCases {
		if err := req.UnmarshalText([]byte(c)); err == nil {
			t.Errorf("%q: expected error, got nil", c)
		}
	}
}
