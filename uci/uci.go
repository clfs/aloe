package uci

import (
	"github.com/clfs/aloe/engine"
)

type Adapter struct {
	e *engine.Engine
}

func NewAdapter(e *engine.Engine) *Adapter {
	return &Adapter{e}
}

func (a *Adapter) SendLine(s string) ([]Response, error) {
	if s == "" {
		return nil, nil
	}

	req, err := toRequest(s)
	if err != nil {
		return nil, err
	}

	return a.Send(req)
}

func (a *Adapter) Send(req Request) ([]Response, error) {
	switch req := req.(type) {
	default:
		return nil, ErrUnknownRequest{req}
	case *RequestUCI:
		return a.handleUCI(*req)
	case *RequestIsReady:
		return a.handleIsReady(*req)
	case *RequestUCINewGame:
		return a.handleUCINewGame(*req)
	case *RequestPosition:
		return a.handlePosition(*req)
	}
}

func (a *Adapter) handleUCI(req RequestUCI) ([]Response, error) {
	return []Response{
		&ResponseID{Name: a.e.Name()},
		&ResponseID{Author: a.e.Author()},
		&ResponseUCIOk{},
	}, nil
}

func (a *Adapter) handleIsReady(req RequestIsReady) ([]Response, error) {
	return []Response{
		&ResponseReadyOk{},
	}, nil
}

func (a *Adapter) handleUCINewGame(req RequestUCINewGame) ([]Response, error) {
	return nil, a.e.NewGame()
}

func (a *Adapter) handlePosition(req RequestPosition) ([]Response, error) {
	if err := a.e.NewGameFromFEN(req.FEN); err != nil {
		return nil, err
	}

	for _, move := range req.Moves {
		if err := a.e.MoveAlgebraic(move); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
