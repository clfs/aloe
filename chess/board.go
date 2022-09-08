package chess

type Board struct {
	// Bitboards for pieces by color.
	White, Black Bitboard
	// Bitboards for pieces by role.
	Pawns, Knights, Bishops, Rooks, Queens, Kings Bitboard
}

// NewBoard returns a new board with all pieces in their starting positions.
func NewBoard() Board {
	return Board{
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

// KingOf returns the square of the king of the given color.
func (b *Board) KingOf(c Color) Square {
	bb := b.ByColor(c) | b.Kings
	return bb.Square()
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
		b.White.Set(s)
	default: // Black
		b.Black.Set(s)
	}

	switch p.Role {
	case Pawn:
		b.Pawns.Set(s)
	case Knight:
		b.Knights.Set(s)
	case Bishop:
		b.Bishops.Set(s)
	case Rook:
		b.Rooks.Set(s)
	case Queen:
		b.Queens.Set(s)
	default: // King
		b.Kings.Set(s)
	}
}

// Remove removes a piece from the square, if any.
func (b *Board) Remove(s Square) {
	b.White.Clear(s)
	b.Black.Clear(s)
	b.Pawns.Clear(s)
	b.Knights.Clear(s)
	b.Bishops.Clear(s)
	b.Rooks.Clear(s)
	b.Queens.Clear(s)
	b.Kings.Clear(s)
}

// At returns the piece on the square, if any.
func (b *Board) At(s Square) (Piece, bool) {
	var p Piece

	switch {
	case b.White.Get(s):
		p.Color = White
	case b.Black.Get(s):
		p.Color = Black
	default:
		return p, false
	}

	switch {
	case b.Pawns.Get(s):
		p.Role = Pawn
	case b.Knights.Get(s):
		p.Role = Knight
	case b.Bishops.Get(s):
		p.Role = Bishop
	case b.Rooks.Get(s):
		p.Role = Rook
	case b.Queens.Get(s):
		p.Role = Queen
	default: // b.kings.Get(s)
		p.Role = King
	}

	return p, true
}
