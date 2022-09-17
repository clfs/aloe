// Package engine implements the Aloe chess engine.
package engine

import (
	"github.com/clfs/aloe/uci"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

// ID returns the engine's UCI identity.
func (e *Engine) ID() uci.ID {
	return uci.ID{
		Name:   "Aloe",
		Author: "Calvin Figuereo-Supraner",
	}
}

// Search runs a search with the provided parameters. The results are encoded in
// a UCI-compatible format. The search is terminated when the channel is closed.
func (e *Engine) Search(p uci.Position, g uci.Go, ch <-chan uci.Info) error {
	return nil // TODO: implement
}
