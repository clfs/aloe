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
