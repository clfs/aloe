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
	req, err := parse(line)
	if err != nil {
		return nil, err
	}
	return a.Send(req)
}

func (a *Adapter) Send(req Request) ([]Response, error) {
	return nil, nil
}

type Request interface{}

type Response interface{}

func parse(line string) (Request, error) {
	return nil, nil
}
