// Package engine implements the Aloe chess engine.
package engine

import (
	"github.com/clfs/aloe/uci"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

// UCIID identifies the engine to the UCI protocol.
func (e *Engine) UCIID() uci.ID {
	return uci.ID{
		Name:   "Aloe",
		Author: "Calvin Figuereo-Supraner",
	}
}

// UCIGo runs a search with the provided parameters. The results are encoded in
// a UCI-compatible format. The search is terminated when the channel is closed.
func (e *Engine) UCIGo(g uci.Go, ch <-chan uci.Info) error {
	return nil // TODO: implement
}
