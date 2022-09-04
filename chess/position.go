package chess

// Position is a chess position.
type Position struct {
	board           Board           // Piece placements.
	sideToMove      Color           // Color of the player to move.
	castlingRights  CastlingRights  // Set of available castling rights.
	enPassantTarget EnPassantTarget // Target square for en passant capture, if any.
	plySinceStart   uint16          // Number of plies since the start of the game.
	ply50MoveRule   uint8           // Number of plies since the last capture or pawn move.
}

// SideToMove returns the color of the player to move.
func (p *Position) SideToMove() Color {
	return p.sideToMove
}

// Board returns the board.
func (p *Position) Board() Board {
	return p.board
}

// LegalMoves returns a list of legal moves.
func (p *Position) LegalMoves() []Move {
	return nil
}
