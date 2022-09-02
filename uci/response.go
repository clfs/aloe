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

type ResponseReadyOk struct{}

func (resp *ResponseReadyOk) MarshalText() ([]byte, error) {
	return []byte("readyok"), nil
}

type ResponseBestMove struct {
	Move   string
	Ponder string
}

func (resp *ResponseBestMove) MarshalText() ([]byte, error) {
	switch {
	case resp.Move == "" && resp.Ponder == "":
		return nil, fmt.Errorf("cannot marshal empty bestmove")
	case resp.Move != "" && resp.Ponder != "":
		return []byte(fmt.Sprintf("bestmove %s ponder %s", resp.Move, resp.Ponder)), nil
	case resp.Move != "":
		return []byte(fmt.Sprintf("bestmove %s", resp.Move)), nil
	default:
		return nil, fmt.Errorf("cannot marshal bestmove with empty ponder")
	}
}
