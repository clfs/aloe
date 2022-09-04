package chess

// Position is a chess position.
type Position struct {
	board           Board
	active          Color
	castlingRights  CastlingRights
	enPassantTarget EnPassantTarget
	plySinceStart   uint16
	ply50MoveRule   uint8
}

// Active returns the color of the active player.
func (p *Position) Active() Color {
	return p.active
}

// Board returns the board.
func (p *Position) Board() Board {
	return p.board
}
