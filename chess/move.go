package chess

// Move is a chess move. It contains information about the starting square,
// destination square, and promotion piece (if any).
//
// Special cases:
//   - The zero value for Move is a null move.
//   - Castling moves use the starting and destination squares of the king.
type Move uint16

// Move bit-packing constants.
const (
	moveFromMask       = 0x003f
	moveFromShift      = 0
	moveToMask         = 0x0fc0
	moveToShift        = 6
	movePromotionMask  = 0x7000
	movePromotionShift = 12
)

// From returns the starting square of a move. For castling moves, From returns
// the starting square of the king.
func (m Move) From() Square {
	return Square(m >> moveFromShift & moveFromMask)
}

// To returns the destination square of a move. For castling moves, To returns
// the destination square of the king.
func (m Move) To() Square {
	return Square(m >> moveToShift & moveToMask)
}

// Promotion returns promotion information for a move. If the move is not a
// promotion, NoPromotion is returned.
func (m Move) Promotion() Promotion {
	return Promotion(m >> movePromotionShift & movePromotionMask)
}
