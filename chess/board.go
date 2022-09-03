package chess

type Board struct {
	White   Bitboard
	Black   Bitboard
	Pawns   Bitboard
	Knights Bitboard
	Bishops Bitboard
	Rooks   Bitboard
	Queens  Bitboard
	Kings   Bitboard
}

func (b *Board) ByRole(r Role) Bitboard {
	switch r {
	case Pawn:
		return b.Pawns
	case Knight:
		return b.Knights
	case Bishop:
		return b.Bishops
	case Rook:
		return b.Rooks
	case Queen:
		return b.Queens
	default: // King
		return b.Kings
	}
}

// ByColor returns a bitboard for pieces matching the given color.
func (b *Board) ByColor(c Color) Bitboard {
	if c == White {
		return b.White
	}
	return b.Black
}
