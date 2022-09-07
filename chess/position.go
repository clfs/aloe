package chess

// Position is a chess position.
type Position struct {
	board        Board
	sideToMove   Color
	castleRights CastleRights

	enPassantFlag   bool   // Whether the previous move was a double pawn push.
	enPassantTarget Square // The en passant target square. Only valid if enPassantFlag is true.

	plySinceStart uint16 // plies since start of the game
	ply50MoveRule uint8  // plies since last capture or pawn move
}

// NewPosition returns a new starting position.
func NewPosition() Position {
	return Position{
		board:        NewBoard(),
		castleRights: NewCastleRights(),
	}
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

// CastleRights returns both players' castle rights.
func (p *Position) CastleRights() CastleRights {
	return p.castleRights
}

// Board returns the board.
func (p *Position) Board() Board {
	return p.board
}

// EnPassantInfo returns information about the en passant target square.
func (p *Position) EnPassantInfo() (Square, bool) {
	return p.enPassantTarget, p.enPassantFlag
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
	// Move the piece back.

	// Restore the captured piece, if any.
	if u.WasCapture {
		capturedPiece := Piece{p.sideToMove, u.CapturedRole}
		p.board.Put(capturedPiece, u.Move.To)
	}

	// Restore en passant settings.
	p.enPassantFlag = u.EnPassantFlag
	p.enPassantTarget = u.EnPassantTarget

	// Restore castling rights.
	p.castleRights = u.CastleRights
}

// HalfMoveClock returns the number of half moves since the last capture or pawn move.
func (p *Position) HalfMoveClock() uint8 {
	return p.ply50MoveRule
}

// FullMoveNumber returns the number of full moves since the start of the game.
func (p *Position) FullMoveNumber() uint16 {
	return p.plySinceStart/2 + 1
}
