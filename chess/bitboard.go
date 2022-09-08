package chess

import "math/bits"

// Bitboard is a bitset of squares.
// From LSB to MSB, the bits represent A1, B1, ..., H1, A2, ..., G8, H8.
type Bitboard uint64

// Reset sets all bits to 0.
func (b *Bitboard) Reset() {
	*b = 0
}

// Set sets the bit at the given square.
func (b *Bitboard) Set(s Square) {
	*b |= 1 << s
}

// Get returns true if the bit at the given square is set.
func (b *Bitboard) Get(s Square) bool {
	return *b&(1<<s) != 0
}

// Clear clears the bit at the given square.
func (b *Bitboard) Clear(s Square) {
	*b &= ^(1 << s)
}

// Toggle toggles the bit at the given square.
func (b *Bitboard) Toggle(s Square) {
	*b ^= 1 << s
}

// IsEmpty returns true if all bits are 0.
func (b *Bitboard) IsEmpty() bool {
	return *b == 0
}

// IsFull returns true if all bits are 1.
func (b *Bitboard) IsFull() bool {
	return *b == ^Bitboard(0)
}

// Square returns the lowest set square. It is invalid to call Square on an
// empty bitboard.
func (b *Bitboard) Square() Square {
	return Square(bits.TrailingZeros64(uint64(*b)))
}
