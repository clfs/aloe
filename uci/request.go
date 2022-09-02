package uci

import (
	"encoding"
	"fmt"
	"strings"
	"time"
)

type UnknownRequestError struct {
	Request Request
}

func (e UnknownRequestError) Error() string {
	return fmt.Sprintf("unknown request: %s", e.Request)
}

func toRequest(s string) (Request, error) {
	words := strings.Fields(s)
	if len(words) == 0 {
		return nil, fmt.Errorf("empty request")
	}

	switch command := words[0]; command {
	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	case "uci":
		return &RequestUCI{}, nil
	}
}

type Request interface {
	encoding.TextUnmarshaler
}

type RequestUCI struct{}

func (req *RequestUCI) UnmarshalText(text []byte) error {
	_ = text
	return nil
}

type RequestIsReady struct{}

func (req *RequestIsReady) UnmarshalText(text []byte) error {
	_ = text
	return nil
}

type RequestUCINewGame struct{}

func (req *RequestUCINewGame) UnmarshalText(text []byte) error {
	_ = text
	return nil
}

type RequestPosition struct {
	FEN   string
	Moves []string
}

const startPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func (req *RequestPosition) UnmarshalText(text []byte) error {
	words := strings.Fields(string(text))
	if len(words) < 2 {
		return fmt.Errorf("invalid position command: %s", text)
	}

	if words[1] == "startpos" {
		req.FEN = startPos
		if len(words) > 2 {
			req.Moves = words[3:]
		}
		return nil
	}

	req.FEN = strings.Join(words[1:7], " ")
	if len(words) > 7 {
		req.Moves = words[8:]
	}
	return nil
}

// RequestGo represents the "go" command.
// Zero values for non-bool fields represent absent controls and must be ignored.
// Time controls are rounded down to the millisecond when sent to the engine.
type RequestGo struct {
	// Search controls.
	SearchMoves []string // SearchMoves restricts the search to the given moves.
	Infinite    bool     // Infinite is true if the engine should search until the "stop" command.
	Ponder      bool     // Ponder is true if the engine is allowed to ponder during the search.
	Depth       int      // Depth is the maximum search depth in plies.
	Nodes       int      // Nodes is the maximum number of nodes to search.
	Mate        int      // Mate is the maximum number of moves to search for a mate.

	// Time controls.
	WhiteTimeRemaining time.Duration // WhiteTimeRemaining is the amount of time remaining for White.
	BlackTimeRemaining time.Duration // BlackTimeRemaining is the amount of time remaining for Black.
	WhiteIncrement     time.Duration // WhiteIncrement is the amount of time added to White's clock after each move.
	BlackIncrement     time.Duration // BlackIncrement is the amount of time added to Black's clock after each move.
	MoveTime           time.Duration // MoveTime is the maximum amount of time to search for a move.
	MovesToGo          int           // MovesToGo is the number of moves until the next time control.
}

func (req *RequestGo) UnmarshalText(text []byte) error {
	_ = text
	return nil // TODO: implement.
}

type RequestStop struct{}

func (req *RequestStop) UnmarshalText(text []byte) error {
	_ = text
	return nil
}

type RequestPonderHit struct{}

func (req *RequestPonderHit) UnmarshalText(text []byte) error {
	_ = text
	return nil
}

type RequestQuit struct{}

func (req *RequestQuit) UnmarshalText(text []byte) error {
	_ = text
	return nil
}
