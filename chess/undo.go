package chess

type Undo struct {
	Move Move // The move that was made.

	WasCapture   bool // Was the move a capture?
	CapturedRole Role // Role of the captured piece, if any.

	EnPassantPresent bool   // Previous value of [Position.enPassant].
	EnPassantTarget  Square // Previous value of [Position.enPassantTarget].

	CastlingRights CastlingRights // Previous value of [Position.castlingRights].
}
