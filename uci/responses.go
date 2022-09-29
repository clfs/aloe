package uci

import (
	"encoding"
	"fmt"
)

// A Response is a command sent from the engine to the client.
type Response interface {
	encoding.TextMarshaler
}

// BestMove represents the "bestmove" UCI command.
type BestMove struct {
	Move   string
	Ponder string
}

func (b BestMove) MarshalText() ([]byte, error) {
	var res []byte

	if b.Move == "" {
		return nil, fmt.Errorf("invalid bestmove: move is empty")
	}

	res = fmt.Appendf(res, "bestmove %s", b.Move)
	if b.Ponder != "" {
		res = fmt.Appendf(res, " ponder %s", b.Ponder)
	}

	return res, nil
}

// ID represents the "id" command.
type ID struct {
	Name   string
	Author string
}

func (i ID) MarshalText() ([]byte, error) {
	var res []byte

	if i.Name == "" {
		return nil, fmt.Errorf("invalid id: name is empty")
	}

	res = fmt.Appendf(res, "id name %s\n", i.Name)

	if i.Author == "" {
		return nil, fmt.Errorf("invalid id: author is empty")
	}

	res = fmt.Appendf(res, "id author %s", i.Author)

	return res, nil
}

// Score types used in [Info].
const (
	ScoreTypeCentipawn = "cp"
	ScoreTypeMate      = "mate"
)

// Info represents the "info" UCI command.
type Info struct {
	Depth     int      // Search depth in plies.
	PV        []string // Moves in the principal variation.
	Score     int      // Score from the engine's point of view.
	ScoreType string   // Either ScoreTypeCentipawn or ScoreTypeMate.
}
