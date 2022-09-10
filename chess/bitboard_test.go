package chess

import (
	"fmt"
)

func ExampleBitboard_Square() {
	var b Bitboard

	b.Set(G1)
	fmt.Println(b.Square() == G1)

	b.Set(A7)
	fmt.Println(b.Square() == G1)
	// Output:
	// true
	// true
}
