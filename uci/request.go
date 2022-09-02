package uci

import (
	"encoding"
	"fmt"
	"strings"
)

type ErrUnknownRequest struct {
	Request Request
}

func (e ErrUnknownRequest) Error() string {
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
