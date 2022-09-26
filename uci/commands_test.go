package uci

import (
	"bytes"
	"testing"
)

func TestBestMove_MarshalText(t *testing.T) {
	cases := []struct {
		in      BestMove
		want    []byte
		wantErr bool
	}{
		{
			in:   BestMove{Move: "e2e4"},
			want: []byte("bestmove e2e4"),
		},
		{
			in:   BestMove{Move: "e2e4", Ponder: "e7e5"},
			want: []byte("bestmove e2e4 ponder e7e5"),
		},
		{
			in:      BestMove{Ponder: "e7e5"},
			wantErr: true,
		},
	}

	for _, c := range cases {
		got, err := c.in.MarshalText()
		if c.wantErr != (err != nil) {
			t.Errorf("%+v: wantErr = %t, err = %v", c.in, c.wantErr, err)
		}
		if !bytes.Equal(c.want, got) {
			t.Errorf("%+v: want = %q, got = %q", c.in, c.want, got)
		}
	}
}
