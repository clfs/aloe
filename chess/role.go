package chess

// Role is the role of a piece.
type Role uint8

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (r Role) IsValid() bool {
	return r <= King
}
