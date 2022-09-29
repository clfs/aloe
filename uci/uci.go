// Package uci describes the Universal Chess Interface (UCI) protocol.
package uci

import "errors"

type Engine interface {
	// Do blocks until the request is completed or an error occurs.
	Do(req Request) error

	// Respond returns and consumes the oldest response from the engine.
	// It blocks until a response is available or an error occurs.
	Respond() (Response, error)

	// Close closes the engine. Any blocked Do or Respond calls will be
	// unblocked and return ErrEngineClosed.
	Close() error
}

var ErrEngineClosed = errors.New("engine closed")
