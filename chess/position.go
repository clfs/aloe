package chess

// Position is a chess position.
type Position struct {
	board          Board
	turn           Color
	castlingRights CastlingRights
	enPassant      struct {
		square Square
		ok     bool
	}
}
