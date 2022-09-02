package chess

type Piece int8

const (
	NoPiece Piece = iota - 1
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	_
	_
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func (p Piece) Color() Color {
	return Color(p >> 3)
}

func (p Piece) Role() Role {
	return Role(p & 7)
}
