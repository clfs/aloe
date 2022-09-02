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
