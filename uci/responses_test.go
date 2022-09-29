package uci

import (
	"bytes"
	"encoding"
	"testing"
)

var testsMarshalResponse = []struct {
	in      encoding.TextMarshaler
	want    []byte
	wantErr bool
}{
	{in: ResponseBestMove{Move: "e2e4"}, want: []byte("bestmove e2e4")},
	{in: ResponseBestMove{Move: "e2e4", Ponder: "e7e5"}, want: []byte("bestmove e2e4 ponder e7e5")},
	{in: ResponseBestMove{Ponder: "e7e5"}, wantErr: true},
	{in: ResponseID{Name: "Skynet", Author: "Cyberdyne"}, want: []byte("id name Skynet\nid author Cyberdyne")},
	{in: ResponseID{Name: "Skynet"}, wantErr: true},
	{in: ResponseID{Author: "Cyberdyne"}, wantErr: true},
}

func TestMarshalResponse(t *testing.T) {
	for _, tc := range testsMarshalResponse {
		got, err := tc.in.MarshalText()
		if tc.wantErr != (err != nil) {
			t.Errorf("%#v: wantErr = %t, err = %v", tc.in, tc.wantErr, err)
		}
		if !bytes.Equal(tc.want, got) {
			t.Errorf("%#v: want = %q, got = %q", tc.in, tc.want, got)
		}
	}
}
