package chess

// Position is a chess position.
type Position struct {
	board          Board          // Piece placements.
	sideToMove     Color          // Color of the player to move.
	castlingRights CastlingRights // Available castling rights.

	enPassantPresent bool   // True if an en passant square is present (legal or not).
	enPassantTarget  Square // Capture square for en passant, if any.

	plySinceStart uint16 // Number of plies since the start of the game.
	ply50MoveRule uint8  // Number of plies since the last capture or pawn move.
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

// IsLegalMove returns true if the move is legal in the position. It does not
// account for insufficient material.
func (p *Position) IsLegalMove(m Move) bool {
	return false
}

// Move updates the position by making a move. It returns an [Undo] that can be
// used to undo the move.
//
// The move must be legal in the position. If not, behavior is undefined.
func (p *Position) Move(m Move) *Undo { return nil }

func (p *Position) Undo(u *Undo) {

}
