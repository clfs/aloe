package fen

import (
	"testing"

	"github.com/clfs/aloe/chess"
	"github.com/google/go-cmp/cmp"
)

func TestEncode(t *testing.T) {
	pos := chess.NewPosition()
	want := StartingFEN

	got, err := Encode(pos)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

func TestDecode(t *testing.T) {
	fen := StartingFEN
	want := chess.NewPosition()

	got, err := Decode(fen)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(-want, +got):\n%s", diff)
	}
}

var validFENTests = []string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
	"8/4npk1/5p1p/1Q5P/1p4P1/4r3/7q/3K1R2 b - - 1 49",
	"5r1k/6pp/4Qpb1/p7/8/6PP/P4PK1/3q4 b - - 4 37",
	"8/8/2P5/4B3/1Q6/4K3/6P1/3k4 w - - 5 67",
	"r2q1rk1/pp2ppbp/2p2np1/6B1/3PP1b1/Q1P2N2/P4PPP/3RKB1R b K - 0 13",
}

var invalidFENTests = []string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/ w KQkq - 0 1",
	" ",
	"",
}

func TestDecode_Valid(t *testing.T) {
	for _, fen := range validFENTests {
		if _, err := Decode(fen); err != nil {
			t.Errorf("failed to decode %q: %v", fen, err)
		}
	}
}

func TestDecode_Invalid(t *testing.T) {
	for _, fen := range invalidFENTests {
		if _, err := Decode(fen); err == nil {
			t.Errorf("decoded invalid FEN %q", fen)
		}
	}
}

func FuzzRoundTrip(f *testing.F) {
	f.Add(StartingFEN)
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
