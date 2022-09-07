package chess

// Undo contains information about a move that can be used to undo it.
type Undo struct {
	Move Move // The move to undo.

	WasCapture   bool // Was the move a capture?
	CapturedRole Role // Role of the captured piece, if any.

	EnPassantFlag   bool   // Was the move to undo preceded by a double pawn push?
	EnPassantSquare Square // En passant square, if any.

	CastleRights CastleRights // Previous castle rights.
}
