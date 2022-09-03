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

// NewBoard returns a new board with all pieces in their starting positions.
func NewBoard() *Board {
	return &Board{
		White:   Bitboard(0x0000_0000_0000_FFFF),
		Black:   Bitboard(0xFFFF_0000_0000_0000),
		Pawns:   Bitboard(0x00FF_0000_0000_FF00),
		Knights: Bitboard(0x4200_0000_0000_0042),
		Bishops: Bitboard(0x2400_0000_0000_0024),
		Rooks:   Bitboard(0x8100_0000_0000_0081),
		Queens:  Bitboard(0x0800_0000_0000_0008),
		Kings:   Bitboard(0x1000_0000_0000_0010),
	}
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

// ByColor returns a bitboard of pieces by color.
func (b *Board) ByColor(c Color) Bitboard {
	if c == White {
		return b.White
	}
	return b.Black
}
