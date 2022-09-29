package uci

import (
	"encoding"
	"fmt"
	"regexp"
	"strings"

	"github.com/clfs/aloe/fen"
)

// A Request is a command sent from the client to the engine.
type Request interface {
	encoding.TextUnmarshaler
}

// Go represents the "go" command.
type Go struct {
	SearchMoves []string // Restrict search to these moves only. Ignore if empty.

	Ponder   bool // Search in pondering mode.
	Infinite bool // Search until interrupted.
	MoveTime int  // If > 0, search for this many milliseconds.

	WhiteTime      int // If > 0, white's remaining time in milliseconds.
	BlackTime      int // If > 0, black's remaining time in milliseconds.
	WhiteIncrement int // If > 0, white's increment in milliseconds.
	BlackIncrement int // If > 0, black's increment in milliseconds.

	Depth     int // If > 0, search this many plies only.
	Nodes     int // If > 0, search this many nodes only.
	Mate      int // If > 0, search for a mate in this many moves.
	MovesToGo int // If > 0, there are this many moves until the next time control.
}

func (g *Go) UnmarshalText(text []byte) error {
	return nil // TODO: implement
}

// Position represents the "position" command.
type Position struct {
	FEN   string
	Moves []string
}

// DefaultPosition is the default [Position] used when no "position" command is
// available.
var DefaultPosition = Position{FEN: fen.StartingFEN}

// Regular expressions for parsing "position" commands.
var (
	rgxPositionStartposMoves = regexp.MustCompile(`^position startpos moves (.+)$`)
	rgxPositionFENMoves      = regexp.MustCompile(`^position fen (.+) moves (.+)$`)
	rgxPositionFEN           = regexp.MustCompile(`^position fen (.+)$`)
)

func (p *Position) UnmarshalText(text []byte) error {
	s := string(text)

	if s == "position startpos" {
		*p = Position{fen.StartingFEN, nil}
		return nil
	}

	// position startpos moves <moves>
	if m := rgxPositionStartposMoves.FindStringSubmatch(s); m != nil {
		*p = Position{fen.StartingFEN, strings.Fields(m[1])}
		return nil
	}

	// position fen <fen> moves <moves>
	if m := rgxPositionFENMoves.FindStringSubmatch(s); m != nil {
		*p = Position{m[1], strings.Fields(m[2])}
		return nil
	}

	// position fen <fen>
	if m := rgxPositionFEN.FindStringSubmatch(s); m != nil {
		*p = Position{m[1], nil}
		return nil
	}

	return fmt.Errorf("invalid position command: %s", text)
}
