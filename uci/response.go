package uci

import (
	"encoding"
	"fmt"
)

type Response interface {
	encoding.TextMarshaler
}

type ResponseID struct {
	Name   string
	Author string
}

func (resp *ResponseID) MarshalText() ([]byte, error) {
	switch {
	case resp.Name != "" && resp.Author != "":
		return nil, fmt.Errorf("cannot marshal both name and author")
	case resp.Name != "":
		return []byte(fmt.Sprintf("id name %s", resp.Name)), nil
	case resp.Author != "":
		return []byte(fmt.Sprintf("id author %s", resp.Author)), nil
	default:
		return nil, fmt.Errorf("cannot marshal empty id")
	}
}

type ResponseUCIOk struct{}

func (resp *ResponseUCIOk) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}
