package uci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	cases := []struct {
		line string
		want Message
	}{
		{
			"uci",
			&UCI{},
		},
		{
			"id name X",
			&IDName{Name: "X"},
		},
		{
			"id author X",
			&IDAuthor{Author: "X"},
		},
	}
	for i, c := range cases {
		got, err := Parse(c.line)
		if err != nil {
			t.Errorf("%d: error: %v", i, err)
		}
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("%d: (-want, +got): %s", i, diff)
		}
	}
}
