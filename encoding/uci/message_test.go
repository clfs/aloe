package uci

import (
	"errors"
	"testing"
)

type marshalTextTest struct {
	in  Message
	out string
	err error
}

var marshalTextTests = []marshalTextTest{
	{
		in:  &BestMove{Move: "e2e4"},
		out: "bestmove e2e4",
	},
	{
		in:  &BestMove{Move: "e2e4", Ponder: "e7e5"},
		out: "bestmove e2e4 ponder e7e5",
	},
	{
		in:  &BestMove{Move: "e2e4", Ponder: "e7e5"},
		out: "bestmove e2e4 ponder e7e5",
	},
}

func TestMarshalText(t *testing.T) {
	for i, tt := range marshalTextTests {
		got, err := tt.in.MarshalText()

		if !errors.Is(err, tt.err) {
			t.Errorf("%d: error %v, want %v", i, err, tt.err)
		}

		if string(got) != tt.out {
			t.Errorf("%d: got %q, want %q", i, got, tt.out)
		}
	}
}
