package chess

// Piece is a chess piece. Bits 0 to 2 represent the role, and bit 3 represents
// the color.
type Piece int8

// NoPiece is a sentinel value for the absence of a piece. Calling Piece methods
// on NoPiece can return invalid results.
const NoPiece Piece = -1

// White pieces.
const (
	WhitePawn Piece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
)

// Black pieces.
const (
	BlackPawn Piece = 8 + iota
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

// Color returns the color of the piece.
func (p Piece) Color() Color {
	return Color(p&8 == 1)
}

// Role returns the role of the piece.
func (p Piece) Role() Role {
	return Role(p & 7)
}
