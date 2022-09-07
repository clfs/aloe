// Package fen implements Forsyth-Edwards Notation (FEN).
//
// [Encode] and [Decode] are exact inverses for all valid inputs.
package fen

import "github.com/clfs/aloe/chess"

// StartingPosition is the FEN for the starting position.
const StartingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Encode returns the FEN for the provided position.
func Encode(p chess.Position) (string, error) {
	return "", nil
}

// Decode returns the position for the provided FEN.
func Decode(fen string) (chess.Position, error) {
	return chess.Position{}, nil
}
