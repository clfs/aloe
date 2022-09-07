// Package fen implements Forsyth-Edwards Notation (FEN).
//
// [Encode] and [Decode] are exact inverses for all accepted inputs.
//
// Accepted inputs must be syntactically correct, but do not have to represent
// legal positions.
package fen

import (
	"fmt"
	"strings"

	"github.com/clfs/aloe/chess"
)

// StartingPosition is the FEN for the starting position.
const StartingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Encode returns the FEN for the provided position.
func Encode(p chess.Position) (string, error) {
	var b strings.Builder

	// Encode the board.

	board := p.Board()

	for rank := chess.Rank8; rank <= chess.Rank8; rank-- {
		empty := 0

		for file := chess.FileA; file <= chess.FileH; file++ {
			piece, ok := board.At(chess.SquareAt(file, rank))
			if !ok {
				empty++
				continue
			}

			if empty > 0 {
				fmt.Fprintf(&b, "%d", empty)
				empty = 0
			}

			s, err := encodePiece(piece)
			if err != nil {
				return "", err
			}

			b.WriteString(s)
		}

		if empty > 0 {
			fmt.Fprintf(&b, "%d", empty)
		}

		if rank > chess.Rank1 {
			b.WriteString("/")
		}
	}

	// Encode the active color.

	b.WriteString(" ")

	s, err := encodeColor(p.SideToMove())
	if err != nil {
		return "", err
	}

	b.WriteString(s)

	// Encode the castle rights.

	b.WriteString(" ")

	s, err = encodeCastleRights(p.CastleRights())
	if err != nil {
		return "", err
	}

	b.WriteString(s)

	// Encode the en passant square.

	b.WriteString(" ")

	s, err = encodeEnPassant(p.EnPassantInfo())
	if err != nil {
		return "", err
	}

	b.WriteString(s)

	// Encode the halfmove clock.

	b.WriteString(" ")

	fmt.Fprintf(&b, "%d", p.HalfMoveClock())

	// Encode the fullmove number.

	b.WriteString(" ")

	fmt.Fprintf(&b, "%d", p.FullMoveNumber())

	return b.String(), nil
}

// Decode returns the position for the provided FEN.
func Decode(fen string) (chess.Position, error) {
	var pos chess.Position

	return pos, nil
}

var stringToPiece = map[string]chess.Piece{
	"P": {Color: chess.White, Role: chess.Pawn},
	"N": {Color: chess.White, Role: chess.Knight},
	"B": {Color: chess.White, Role: chess.Bishop},
	"R": {Color: chess.White, Role: chess.Rook},
	"Q": {Color: chess.White, Role: chess.Queen},
	"K": {Color: chess.White, Role: chess.King},
	"p": {Color: chess.Black, Role: chess.Pawn},
	"n": {Color: chess.Black, Role: chess.Knight},
	"b": {Color: chess.Black, Role: chess.Bishop},
	"r": {Color: chess.Black, Role: chess.Rook},
	"q": {Color: chess.Black, Role: chess.Queen},
	"k": {Color: chess.Black, Role: chess.King},
}

var stringToColor = map[string]chess.Color{
	"w": chess.White,
	"b": chess.Black,
}

var stringToCastleRights = map[string]chess.CastleRights{
	"-":    0,
	"K":    chess.WhiteOO,
	"Q":    chess.WhiteOOO,
	"k":    chess.BlackOO,
	"q":    chess.BlackOOO,
	"KQ":   chess.WhiteOO | chess.WhiteOOO,
	"Kk":   chess.WhiteOO | chess.BlackOO,
	"Kq":   chess.WhiteOO | chess.BlackOOO,
	"Qk":   chess.WhiteOOO | chess.BlackOO,
	"Qq":   chess.WhiteOOO | chess.BlackOOO,
	"kq":   chess.BlackOO | chess.BlackOOO,
	"KQk":  chess.WhiteOO | chess.WhiteOOO | chess.BlackOO,
	"KQq":  chess.WhiteOO | chess.WhiteOOO | chess.BlackOOO,
	"Kkq":  chess.WhiteOO | chess.BlackOO | chess.BlackOOO,
	"Qkq":  chess.WhiteOOO | chess.BlackOO | chess.BlackOOO,
	"KQkq": chess.WhiteOO | chess.WhiteOOO | chess.BlackOO | chess.BlackOOO,
}

var stringToSquare = map[string]chess.Square{
	"a3": chess.A3,
	"a6": chess.A6,
	"b3": chess.B3,
	"b6": chess.B6,
	"c3": chess.C3,
	"c6": chess.C6,
	"d3": chess.D3,
	"d6": chess.D6,
	"e3": chess.E3,
	"e6": chess.E6,
	"f3": chess.F3,
	"f6": chess.F6,
	"g3": chess.G3,
	"g6": chess.G6,
	"h3": chess.H3,
	"h6": chess.H6,
}

// inverseLookup returns the key for the provided value, if any. O(n), but that's fine.
func inverseLookup[T comparable](m map[string]T, v T) (string, bool) {
	for k, val := range m {
		if val == v {
			return k, true
		}
	}
	return "", false
}

func decodeColor(s string) (chess.Color, error) {
	c, ok := stringToColor[s]
	if !ok {
		return false, fmt.Errorf("invalid color: %s", s)
	}
	return c, nil
}

func encodeColor(c chess.Color) (string, error) {
	s, ok := inverseLookup(stringToColor, c)
	if !ok {
		return "", fmt.Errorf("invalid color: %v", c)
	}
	return s, nil
}

func decodePiece(s string) (chess.Piece, error) {
	p, ok := stringToPiece[s]
	if !ok {
		return chess.Piece{}, fmt.Errorf("invalid piece: %s", s)
	}
	return p, nil
}

func encodePiece(p chess.Piece) (string, error) {
	s, ok := inverseLookup(stringToPiece, p)
	if !ok {
		return "", fmt.Errorf("invalid piece: %v", p)
	}
	return s, nil
}

func decodeCastleRights(s string) (chess.CastleRights, error) {
	c, ok := stringToCastleRights[s]
	if !ok {
		return 0, fmt.Errorf("invalid castle rights: %s", s)
	}
	return c, nil
}

func encodeCastleRights(c chess.CastleRights) (string, error) {
	s, ok := inverseLookup(stringToCastleRights, c)
	if !ok {
		return "", fmt.Errorf("invalid castle rights: %v", c)
	}
	return s, nil
}

func decodeEnPassant(s string) (chess.Square, bool, error) {
	if s == "-" {
		return 0, false, nil
	}

	sq, ok := stringToSquare[s]
	if !ok {
		return 0, false, fmt.Errorf("invalid en passant square: %s", s)
	}

	return sq, true, nil
}

func encodeEnPassant(sq chess.Square, ok bool) (string, error) {
	if !ok {
		return "-", nil
	}

	s, ok := inverseLookup(stringToSquare, sq)
	if !ok {
		return "", fmt.Errorf("invalid en passant square: %v", sq)
	}

	return s, nil
}
