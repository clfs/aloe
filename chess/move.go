package chess

// Move is a chess move. It contains information about the source square,
// destination square, and promotion piece (if any).
//
// Special cases:
// - The zero value for Move is a null move.
// - Castling moves use the source and destination squares of the king.
type Move uint16

// Move bit-packing constants.
const (
	moveSrcMask        = 0x003f
	moveSrcShift       = 0
	moveDstMask        = 0x0fc0
	moveDstShift       = 6
	movePromotionMask  = 0x7000
	movePromotionShift = 12
)

// Src returns the source square of a move. For castling moves, Src returns the
// starting square of the king.
func (m Move) Src() Square {
	return Square(m >> moveSrcShift & moveSrcMask)
}

// Dst returns the destination square of a move. For castling moves, Dst returns
// the destination square of the king.
func (m Move) Dst() Square {
	return Square(m >> moveDstShift & moveDstMask)
}

// Promotion returns promotion information for a move. If the move is not a
// promotion, NoPromotion is returned.
func (m Move) Promotion() Promotion {
	return Promotion(m >> movePromotionShift & movePromotionMask)
}
