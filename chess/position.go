package chess

// Position is a chess position.
type Position struct {
	board          Board
	sideToMove     Color
	castlingRights CastlingRights

	enPassantFlag   bool   // Whether the previous move was a double pawn push.
	enPassantTarget Square // The en passant target square. Only valid if enPassantFlag is true.

	plySinceStart uint16 // plies since start of the game
	ply50MoveRule uint8  // plies since last capture or pawn move
}

func NewPosition() *Position {
	return nil
}

func (p *Position) MarshalText() ([]byte, error) {
	return nil, nil
}

func (p *Position) UnmarshalText(text []byte) error {
	return nil
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
// account for insufficient material or three-fold repetition.
func (p *Position) IsLegalMove(m Move) bool {
	return false
}

// Move updates the position by making a move. It returns information that can
// be used to undo the move.
//
// The move must be legal by the definition of [Position.IsLegalMove]. If not,
// behavior is undefined.
func (p *Position) Move(m Move) *Undo { return nil }

// Undo undoes a [Position.Move] call.
func (p *Position) Undo(u *Undo) {

}
