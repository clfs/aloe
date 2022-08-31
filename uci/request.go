package uci

import (
	"bytes"
	"encoding"
	"fmt"
)

type Request interface {
	encoding.TextUnmarshaler
}

type ErrInvalidRequest struct{ text []byte }

func (e ErrInvalidRequest) Error() string {
	return fmt.Sprintf("invalid request: %s", e.text)
}

type UCI struct{}

func (u *UCI) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("uci")) {
		return ErrInvalidRequest{text}
	}
	return nil
}

func parse(line string) (Request, error) {
	var req Request

	if line == "uci" {
		req = &UCI{}
		err := req.UnmarshalText([]byte(line))
		return req, err
	}

	return nil, ErrInvalidRequest{[]byte(line)}
}
