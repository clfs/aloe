package fen

import (
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
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
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
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
	for _, fen := range validFENTests {
		f.Add(fen)
	}
	for _, fen := range invalidFENTests {
		f.Add(fen)
	}
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

func FuzzEncodeThenDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pos chess.Position

		c := fuzz.NewConsumer(data)
		if err := c.GenerateStruct(&pos); err != nil {
			return
		}

		fen, err := Encode(pos)
		if err != nil {
			t.Skip() // Not encodable.
		}

		pos2, err := Decode(fen)
		if err != nil {
			t.Errorf("failed to decode %q: %v", fen, err)
		}

		if diff := cmp.Diff(pos, pos2); diff != "" {
			t.Errorf("changed after round trip: (-old, +new):\n%s", diff)
		}
	})
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
	"2B5/1P1r2k1/1P3R1p/1p1P4/8/PpP2p2/P1n4K/8 w - - 80 99",
	"4n2k/2P5/P7/5pp1/2P3P1/RQ6/3ppq1p/KR6 b Kq h6 1 100",
	"8/6P1/nP1B1Ppp/PP1BRN2/PbP1K1RN/pQP2pqp/ppp2r2/2r1nb1k w - - 0 1",
}

var invalidFENTests = []string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/ w KQkq - 0 1",
	" ",
	"",
}
