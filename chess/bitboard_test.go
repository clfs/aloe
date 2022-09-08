package chess

import (
	"fmt"
)

func ExampleBitboard_Reset() {
	b := Bitboard(0xcafe)
	b.Reset()
	fmt.Printf("%#x\n", b)
	// Output:
	// 0x0
}

func ExampleBitboard_Set() {
	var b Bitboard
	b.Set(A3)
	fmt.Printf("%#x\n", b)
	// Output:
	// 0x10000
}

func ExampleBitboard_Get() {
	var b Bitboard
	if b.Get(C6) {
		fmt.Println("set")
	} else {
		fmt.Println("not set")
	}
	// Output:
	// not set
}

func ExampleBitboard_Clear() {
	b := Bitboard(0xcafe)
	b.Clear(B1)
	fmt.Printf("%#x\n", b)
	// Output:
	// 0xcafc
}

func ExampleBitboard_Toggle() {
	b := Bitboard(0xcafe)
	b.Toggle(H2)
	fmt.Printf("%#x\n", b)
	// Output:
	// 0x4afe
}

func ExampleBitboard_IsEmpty() {
	var b Bitboard
	fmt.Println(b.IsEmpty())
	// Output:
	// true
}

func ExampleBitboard_IsFull() {
	b := Bitboard(0xffffffffffffffff)
	fmt.Println(b.IsFull())
	// Output:
	// true
}

func ExampleBitboard_Square() {
	b := Bitboard(0x8)
	fmt.Println(b.Square() == D1)

	b = Bitboard(0xfff8)
	fmt.Println(b.Square() == D1)
	// Output:
	// true
	// true
}
