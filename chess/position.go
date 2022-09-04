package chess

// Position is a chess position.
type Position struct {
	board           Board
	sideToMove      Color
	castlingRights  CastlingRights
	enPassantTarget EnPassantTarget
	plySinceStart   uint16
	ply50MoveRule   uint8
}
