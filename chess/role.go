package chess

import "fmt"

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

func (r Role) String() string {
	switch r {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("Role(%d)", r)
	}
}
