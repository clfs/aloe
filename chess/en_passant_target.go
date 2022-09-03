package chess

// EnPassantTarget represents a target square when capturing en passant.
// The zero value indicates that no target exists.
type EnPassantTarget uint8

// Square returns the en passant target square. If no target square exists, ok is false.
func (e *EnPassantTarget) Square() (s Square, ok bool) {
	return Square(*e - 1), *e > 0
}

// Set sets the en passant target square.
func (e *EnPassantTarget) Set(s Square) {
	*e = EnPassantTarget(s + 1)
}

// Clear discards the current en passant target square.
func (e *EnPassantTarget) Clear() {
	*e = 0
}
