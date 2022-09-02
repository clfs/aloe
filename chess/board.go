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

func (b *Board) Pieces(p Piece) Bitboard {
	switch p.Role() {
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
	case King:
		return b.Kings
	default:
		panic("invalid piece")
	}
}

func (b *Board) Color(c Color) Bitboard {
	switch c {
	case White:
		return b.White
	case Black:
		return b.Black
	default:
		panic("invalid color")
	}
}
