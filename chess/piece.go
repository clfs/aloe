package chess

// Piece is a chess piece.
type Piece struct {
	Color Color
	Role  Role
}

// IsValid returns true if the piece is valid.
func (p Piece) IsValid() bool {
	return p.Role.IsValid()
}
