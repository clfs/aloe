package chess

// Role is the role of a piece.
type Role int8

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)
