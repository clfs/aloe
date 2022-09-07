package chess

type Board struct {
	// Bitboards for pieces by color.
	white, black Bitboard
	// Bitboards for pieces by role.
	pawns, knights, bishops, rooks, queens, kings Bitboard
}

// NewBoard returns a new board with all pieces in their starting positions.
func NewBoard() Board {
	return Board{
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

// Put puts a piece on a square. Any piece already on the square is removed.
func (b *Board) Put(p Piece, s Square) {
	b.Remove(s)
	b.PutDangerous(p, s)
}

// PutDangerous puts a piece on a square without checking if a different piece
// already exists there. It is faster than [Board.Put].
func (b *Board) PutDangerous(p Piece, s Square) {
	switch p.Color {
	case White:
		b.white.Set(s)
	default: // Black
		b.black.Set(s)
	}

	switch p.Role {
	case Pawn:
		b.pawns.Set(s)
	case Knight:
		b.knights.Set(s)
	case Bishop:
		b.bishops.Set(s)
	case Rook:
		b.rooks.Set(s)
	case Queen:
		b.queens.Set(s)
	default: // King
		b.kings.Set(s)
	}
}

// Remove removes a piece from the square, if any.
func (b *Board) Remove(s Square) {
	b.white.Clear(s)
	b.black.Clear(s)
	b.pawns.Clear(s)
	b.knights.Clear(s)
	b.bishops.Clear(s)
	b.rooks.Clear(s)
	b.queens.Clear(s)
	b.kings.Clear(s)
}

// At returns the piece on the square, if any.
func (b *Board) At(s Square) (Piece, bool) {
	var p Piece

	switch {
	case b.white.Get(s):
		p.Color = White
	case b.black.Get(s):
		p.Color = Black
	default:
		return p, false
	}

	switch {
	case b.pawns.Get(s):
		p.Role = Pawn
	case b.knights.Get(s):
		p.Role = Knight
	case b.bishops.Get(s):
		p.Role = Bishop
	case b.rooks.Get(s):
		p.Role = Rook
	case b.queens.Get(s):
		p.Role = Queen
	default: // b.kings.Get(s)
		p.Role = King
	}

	return p, true
}
