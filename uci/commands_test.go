package uci

import (
	"bytes"
	"encoding"
	"testing"
)

var testsMarshalText = []struct {
	in      encoding.TextMarshaler
	want    []byte
	wantErr bool
}{
	{in: BestMove{Move: "e2e4"}, want: []byte("bestmove e2e4")},
	{in: BestMove{Move: "e2e4", Ponder: "e7e5"}, want: []byte("bestmove e2e4 ponder e7e5")},
	{in: BestMove{Ponder: "e7e5"}, wantErr: true},
	{in: ID{Name: "Skynet", Author: "Cyberdyne"}, want: []byte("id name Skynet\nid author Cyberdyne")},
	{in: ID{Name: "Skynet"}, wantErr: true},
	{in: ID{Author: "Cyberdyne"}, wantErr: true},
}

func TestMarshalText(t *testing.T) {
	for _, tc := range testsMarshalText {
		got, err := tc.in.MarshalText()
		if tc.wantErr != (err != nil) {
			t.Errorf("%#v: wantErr = %t, err = %v", tc.in, tc.wantErr, err)
		}
		if !bytes.Equal(tc.want, got) {
			t.Errorf("%#v: want = %q, got = %q", tc.in, tc.want, got)
		}
	}
}
