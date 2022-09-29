package uci

import (
	"bytes"
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

// RequestIsReady represents the "isready" command.
type RequestIsReady struct{}

func (req *RequestIsReady) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("isready")) {
		return fmt.Errorf("invalid isready request")
	}
	return nil
}

// RequestGo represents the "go" command.
type RequestGo struct {
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

func (req *RequestGo) UnmarshalText(text []byte) error {
	return nil // TODO: implement
}

// RequestPosition represents the "position" command.
type RequestPosition struct {
	FEN   string
	Moves []string
}

// Regular expressions for parsing "position" commands.
var (
	rgxPositionStartposMoves = regexp.MustCompile(`^position startpos moves (.+)$`)
	rgxPositionFENMoves      = regexp.MustCompile(`^position fen (.+) moves (.+)$`)
	rgxPositionFEN           = regexp.MustCompile(`^position fen (.+)$`)
)

func (req *RequestPosition) UnmarshalText(text []byte) error {
	s := string(text)

	if s == "position startpos" {
		*req = RequestPosition{fen.StartingFEN, nil}
		return nil
	}

	// position startpos moves <moves>
	if m := rgxPositionStartposMoves.FindStringSubmatch(s); m != nil {
		*req = RequestPosition{fen.StartingFEN, strings.Fields(m[1])}
		return nil
	}

	// position fen <fen> moves <moves>
	if m := rgxPositionFENMoves.FindStringSubmatch(s); m != nil {
		*req = RequestPosition{m[1], strings.Fields(m[2])}
		return nil
	}

	// position fen <fen>
	if m := rgxPositionFEN.FindStringSubmatch(s); m != nil {
		*req = RequestPosition{m[1], nil}
		return nil
	}

	return fmt.Errorf("invalid position command: %s", text)
}

// RequestUCI represents the "uci" command.
type RequestUCI struct{}

func (req *RequestUCI) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("uci")) {
		return fmt.Errorf("invalid uci request")
	}
	return nil
}
