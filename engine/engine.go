// Package engine implements the Aloe chess engine.
package engine

import (
	"fmt"

	"github.com/clfs/aloe/uci"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Do(req uci.Request) error {
	return nil // TODO: implement
}

func (e *Engine) Respond() (uci.Response, error) {
	return nil, fmt.Errorf("not implemented") // TODO: implement
}

func (e *Engine) Close() error {
	return nil // TODO: implement
}
