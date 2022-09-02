package uci

import (
	"bytes"
	"encoding"
	"fmt"
)

type ErrUnknownRequest struct {
	Request Request
}

func (e ErrUnknownRequest) Error() string {
	return fmt.Sprintf("unknown request: %s", e.Request)
}

func ToRequest(s string) (Request, error) {
	b := []byte(s)

	var reqUCI RequestUCI
	if err := reqUCI.UnmarshalText(b); err == nil {
		return &reqUCI, nil
	}

	return nil, fmt.Errorf("unknown request: %q", s)
}

type Request interface {
	encoding.TextUnmarshaler
}

type RequestUCI struct{}

func (req *RequestUCI) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("uci")) {
		return fmt.Errorf("cannot unmarshal %q", text)
	}
	return nil
}
