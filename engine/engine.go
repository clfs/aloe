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
func UCIID() uci.ID {
	return uci.ID{
		Name:   "Aloe",
		Author: "Calvin Figuereo-Supraner",
	}
}

// UCIGo runs a search with the provided parameters. The results are encoded in
// a UCI-compatible format.
func UCIGo(g uci.Go, ch <-chan uci.Info) error {
	return nil // TODO: implement
}
