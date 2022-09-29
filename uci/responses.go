package uci

import (
	"encoding"
	"fmt"
)

// A Response is a command sent from the engine to the client.
type Response interface {
	encoding.TextMarshaler
}

// ResponseBestMove represents the "bestmove" command.
type ResponseBestMove struct {
	Move   string
	Ponder string
}

func (resp ResponseBestMove) MarshalText() ([]byte, error) {
	var text []byte

	if resp.Move == "" {
		return nil, fmt.Errorf("invalid bestmove: move is empty")
	}

	text = fmt.Appendf(text, "bestmove %s", resp.Move)
	if resp.Ponder != "" {
		text = fmt.Appendf(text, " ponder %s", resp.Ponder)
	}

	return text, nil
}

// ResponseID represents the "id" command.
type ResponseID struct {
	Name   string
	Author string
}

func (resp ResponseID) MarshalText() ([]byte, error) {
	var text []byte

	if resp.Name == "" {
		return nil, fmt.Errorf("invalid id: name is empty")
	}

	text = fmt.Appendf(text, "id name %s\n", resp.Name)

	if resp.Author == "" {
		return nil, fmt.Errorf("invalid id: author is empty")
	}

	text = fmt.Appendf(text, "id author %s", resp.Author)

	return text, nil
}

// Score types used in [ResponseInfo].
const (
	ScoreTypeCentipawn = "cp"
	ScoreTypeMate      = "mate"
)

// ResponseInfo represents the "info" UCI command.
type ResponseInfo struct {
	Depth     int      // Search depth in plies.
	PV        []string // Moves in the principal variation.
	Score     int      // Score from the engine's point of view.
	ScoreType string   // Either ScoreTypeCentipawn or ScoreTypeMate.
}
