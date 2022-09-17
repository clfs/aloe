package uci

import (
	"testing"

	"github.com/clfs/aloe/fen"
	"github.com/google/go-cmp/cmp"
)

func TestPosition_UnmarshalText(t *testing.T) {
	cases := []struct {
		in      []byte
		want    Position
		wantErr bool
	}{
		{
			in:   []byte("position startpos"),
			want: Position{FEN: fen.StartingFEN},
		},
	}

	for _, c := range cases {
		var got Position
		if err := got.UnmarshalText(c.in); c.wantErr != (err != nil) {
			t.Errorf("%q: wantErr = %t, but error: %v", c.in, c.wantErr, err)
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("%q: (-want, +got)\n%s", c.in, diff)
		}
	}
}

func TestGo_UnmarshalText(t *testing.T) {
	cases := []struct {
		in      []byte
		want    Go
		wantErr bool
	}{
		{
			in:   []byte("go"),
			want: Go{Infinite: true},
		},
	}

	for _, c := range cases {
		var got Go
		if err := got.UnmarshalText(c.in); c.wantErr != (err != nil) {
			t.Errorf("%q: wantErr = %t, but error: %v", c.in, c.wantErr, err)
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("%q: (-want, +got)\n%s", c.in, diff)
		}
	}
}
