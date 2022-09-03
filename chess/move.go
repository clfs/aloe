package chess

// Move is a chess move. Bits 0 to 5 are the source square, bits 6 to 11 are the
// destination square, and bits 12 to 14 are the promotion role.
//
// The zero value for Move is a null move.
//
// If the move is not a promotion, the promotion bits are all 0.
//
// If the move is a castling move, the source and destination squares are the
// starting squares of the king and rook respectively.
type Move uint16

func (m Move) Src() Square {
	return Square(m & 63)
}

func (m Move) Dst() Square {
	return Square(m>>6) & 63
}

func (m Move) Promotion() Role {
	return Role(m>>12) & 7
}

func (m Move) IsPromotion() bool {
	return m.Promotion() != 0
}
