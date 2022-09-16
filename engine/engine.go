// Package engine implements the Aloe chess engine.
package engine

import (
	"context"

	"github.com/clfs/aloe/uci"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

// UCIID identifies the engine to the UCI protocol.
func UCIID() uci.ID {
	return uci.ID{
		Name:    "Aloe",
		Authors: []string{"Calvin Figuereo-Supraner"},
	}
}

// UCIGo runs a search with the provided parameters. The results are encoded in
// a UCI-compatible format.
func UCIGo(ctx context.Context, g uci.Go) (uci.Info, error) {
	return uci.Info{}, nil // TODO: implement
}
