package chess

type Board struct {
	// Bitboards for pieces by color.
	white, black Bitboard
	// Bitboards for pieces by role.
	pawns, knights, bishops, rooks, queens, kings Bitboard
}

// NewBoard returns a new board with all pieces in their starting positions.
func NewBoard() *Board {
	return &Board{
		white:   Bitboard(0x0000_0000_0000_FFFF),
		black:   Bitboard(0xFFFF_0000_0000_0000),
		pawns:   Bitboard(0x00FF_0000_0000_FF00),
		knights: Bitboard(0x4200_0000_0000_0042),
		bishops: Bitboard(0x2400_0000_0000_0024),
		rooks:   Bitboard(0x8100_0000_0000_0081),
		queens:  Bitboard(0x0800_0000_0000_0008),
		kings:   Bitboard(0x1000_0000_0000_0010),
	}
}

// NewBoardFromFEN returns a new board from a FEN string.
func NewBoardFromFEN(fen string) (*Board, error) {
	return nil, nil
}

func (b *Board) ByRole(r Role) Bitboard {
	switch r {
	case Pawn:
		return b.pawns
	case Knight:
		return b.knights
	case Bishop:
		return b.bishops
	case Rook:
		return b.rooks
	case Queen:
		return b.queens
	default: // King
		return b.kings
	}
}

// ByColor returns a bitboard of pieces by color.
func (b *Board) ByColor(c Color) Bitboard {
	if c == White {
		return b.white
	}
	return b.black
}
