package uci

import (
	"errors"

	"github.com/clfs/aloe/engine"
)

var ErrDone = errors.New("uci: done")

type Adapter struct {
	e    *engine.Engine
	done bool
}

func NewAdapter(e *engine.Engine) *Adapter {
	return &Adapter{e: e, done: false}
}

func (a *Adapter) SendLine(s string) ([]Response, error) {
	req, err := toRequest(s)
	if err != nil {
		return nil, err
	}

	return a.Send(req)
}

func (a *Adapter) Send(req Request) ([]Response, error) {
	if a.done {
		return nil, ErrDone
	}

	switch req := req.(type) {
	default:
		return nil, UnknownRequestError{req}
	case *RequestUCI:
		return a.handleUCI(*req)
	case *RequestIsReady:
		return a.handleIsReady(*req)
	case *RequestUCINewGame:
		return a.handleUCINewGame(*req)
	case *RequestPosition:
		return a.handlePosition(*req)
	case *RequestGo:
		return a.handleGo(*req)
	case *RequestStop:
		return a.handleStop(*req)
	case *RequestPonderHit:
		return a.handlePonderHit(*req)
	case *RequestQuit:
		return a.handleQuit(*req)
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

func (a *Adapter) handleGo(req RequestGo) ([]Response, error) {
	return nil, nil // TODO: implement.
}

func (a *Adapter) handleStop(req RequestStop) ([]Response, error) {
	return nil, nil // TODO: implement.
}

func (a *Adapter) handlePonderHit(req RequestPonderHit) ([]Response, error) {
	return nil, nil // TODO: implement.
}

func (a *Adapter) handleQuit(req RequestQuit) ([]Response, error) {
	return nil, ErrDone
}
