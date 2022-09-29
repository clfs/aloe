package uci

import (
	"testing"

	"github.com/clfs/aloe/fen"
	"github.com/google/go-cmp/cmp"
)

var testsValidPositionCommands = []struct {
	in   string
	want RequestPosition
}{
	{
		"position startpos",
		RequestPosition{FEN: fen.StartingFEN},
	},
	{
		"position startpos moves e2e4",
		RequestPosition{fen.StartingFEN, []string{"e2e4"}},
	},
	{
		"position startpos moves e2e4 e7e5",
		RequestPosition{fen.StartingFEN, []string{"e2e4", "e7e5"}},
	},
	{
		"position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		RequestPosition{FEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
	},
	{
		"position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 moves e2e4",
		RequestPosition{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", []string{"e2e4"}},
	},
	{
		"position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 moves e2e4 e7e5",
		RequestPosition{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", []string{"e2e4", "e7e5"}},
	},
}

var testsInvalidPositionCommands = []string{
	"",
	"position",
	"position fen",
	"position startpos moves",
}

func TestPosition_UnmarshalText(t *testing.T) {
	for _, tc := range testsValidPositionCommands {
		var got RequestPosition
		if err := got.UnmarshalText([]byte(tc.in)); err != nil {
			t.Errorf("%q: error: %v", tc.in, err)
		}
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
		}
	}

	for _, tc := range testsInvalidPositionCommands {
		var got RequestPosition
		if err := got.UnmarshalText([]byte(tc)); err == nil {
			t.Errorf("%q: expected error, got nil", tc)
		}
	}
}

var testsValidGoCommands = []struct {
	in   string
	want RequestGo
}{
	{
		"go",
		RequestGo{Infinite: true},
	},
}

var testsInvalidGoCommands = []string{
	"",
}

func TestGo_UnmarshalText(t *testing.T) {
	for _, tc := range testsValidGoCommands {
		var got RequestGo
		if err := got.UnmarshalText([]byte(tc.in)); err != nil {
			t.Errorf("%q: error: %v", tc.in, err)
		}
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
		}
	}

	for _, tc := range testsInvalidGoCommands {
		var got RequestGo
		if err := got.UnmarshalText([]byte(tc)); err == nil {
			t.Errorf("%q: expected error, got nil", tc)
		}
	}
}
