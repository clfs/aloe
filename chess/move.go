package chess

// Move is a chess move. The zero value for Move is a null move.
type Move struct {
	// The starting square. When castling, this is the king's starting square.
	From Square

	// The destination square. When castling, this is the king's destination square.
	To Square

	// Promotion information. For non-promotion moves, this is [NoPromotion].
	PromotionInfo PromotionInfo
}
