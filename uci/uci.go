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

func (a *Adapter) SendLine(line string) ([]Response, error) {
	req, err := ToRequest(line)
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
		return a.sendUCI(*req)
	}
}

func (a *Adapter) sendUCI(req RequestUCI) ([]Response, error) {
	return []Response{
		&ResponseID{Name: a.e.Name()},
		&ResponseID{Author: a.e.Author()},
		&ResponseUCIOk{},
	}, nil
}
