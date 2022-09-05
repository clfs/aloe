package chess

type Undo struct {
	Move Move // The move that was made.

	WasCapture   bool // Was the move a capture?
	CapturedRole Role // Role of the captured piece, if any.

	EnPassantFlag   bool
	EnPassantTarget Square
	CastlingRights  CastlingRights
}
