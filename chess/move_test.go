package chess

import (
	"fmt"
	"testing"
)

func TestNewMove(t *testing.T) {
	cases := []struct {
		in      string
		want    Move
		wantErr bool
	}{
		{in: "e2e4", want: Move{From: E2, To: E4}},
		{in: "e7e8q", want: Move{From: E7, To: E8, PromotionInfo: QueenPromotion}},
		{in: "a1h8", want: Move{From: A1, To: H8}},
		{in: "8888", wantErr: true},
		{in: "e2e", wantErr: true},
		{in: "e7e8h", wantErr: true},
	}
	for _, c := range cases {
		got, err := NewMove(c.in)

		if (err != nil) != c.wantErr {
			t.Errorf("%q: incorrect error value: %v", c.in, err)
			continue
		}

		if got != c.want {
			t.Errorf("%q: want %v, got %v", c.in, c.want, got)
		}
	}
}

func ExampleNewMove() {
	m, _ := NewMove("e2e4")
	fmt.Printf("%+v\n", m)

	m, _ = NewMove("e7e8q") // promotion
	fmt.Printf("%+v\n", m)

	m, _ = NewMove("e1c1") // white short castling
	fmt.Printf("%+v\n", m)
	// Output:
	// {From:E2 To:E4 PromotionInfo:NoPromotion}
	// {From:E7 To:E8 PromotionInfo:QueenPromotion}
	// {From:E1 To:C1 PromotionInfo:NoPromotion}
}
