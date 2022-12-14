package chess

import "fmt"

// Position is a chess position.
type Position struct {
	Board        Board
	SideToMove   Color
	CastleRights CastleRights

	EnPassantFlag   bool   // Whether the previous move was a double pawn push.
	EnPassantSquare Square // The en passant square. Only valid if enPassantFlag is true.

	FullMoveNumber uint16 // Number of full moves. Starts at 1 and increments after Black moves.
	HalfMoveClock  uint8  // Number of plies since last capture or pawn move.
}

// NewPosition returns a new starting position.
func NewPosition() Position {
	return Position{
		Board:          NewBoard(),
		CastleRights:   NewCastleRights(),
		FullMoveNumber: 1,
	}
}

// LegalMoves returns a list of legal moves.
func (p *Position) LegalMoves() []Move {
	var moves []Move

	return moves
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
		capturedPiece := Piece{p.SideToMove, u.CapturedRole}
		p.Board.Put(capturedPiece, u.Move.To)
	}

	// Restore en passant settings.
	p.EnPassantFlag = u.EnPassantFlag
	p.EnPassantSquare = u.EnPassantSquare

	// Restore castling rights.
	p.CastleRights = u.CastleRights
}

// IsValid returns nil if the position is valid.
func (p *Position) IsValid() error {
	// The underlying board must be valid.
	if err := p.Board.IsValid(); err != nil {
		return err
	}

	// The en passant flag and square must agree.
	if p.EnPassantFlag {
		r := p.EnPassantSquare.Rank()
		if r != 3 && r != 6 {
			return fmt.Errorf("invalid en passant square: %s", p.EnPassantSquare)
		}
	} else {
		if p.EnPassantSquare != 0 {
			return fmt.Errorf("en passant square is set but en passant flag is not set")
		}
	}

	if !p.CastleRights.IsValid() {
		return fmt.Errorf("invalid castling rights")
	}

	return nil
}
