package chess

// EnPassantRight is the right to capture en passant. Bits 0 to 5 are the en
// passant target square (if any), and bit 6 is whether the right is available.
type EnPassantRight uint8

func (e *EnPassantRight) Target() (s Square, ok bool) {
	return Square(*e & 0x3F), *e&(1<<6) != 0
}

func (e *EnPassantRight) SetTarget(s Square) {
	*e = 1<<6 | EnPassantRight(s)
}

func (e *EnPassantRight) Clear() {
	*e = 0
}
