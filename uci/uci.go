package uci

import (
	"fmt"

	"github.com/clfs/aloe/engine"
)

type ErrWrongDirection struct {
	Message Message
}

func (e ErrWrongDirection) Error() string {
	return fmt.Sprintf("wrong direction: %s", e.Message)
}

type Adapter struct {
	e *engine.Engine
}

func NewAdapter(e *engine.Engine) *Adapter {
	return &Adapter{e}
}

func (a *Adapter) SendLine(line string) ([]Message, error) {
	msg, err := Parse(line)
	if err != nil {
		return nil, err
	}

	return a.Send(msg)
}

func (a *Adapter) Send(m Message) ([]Message, error) {
	if !m.EngineBound() {
		return nil, ErrWrongDirection{m}
	}

	switch v := m.(type) {
	default:
		return nil, ErrUnknownMessage{m}
	case *UCI:
		return a.sendUCI(*v)
	}
}

func (a *Adapter) sendUCI(m UCI) ([]Message, error) {
	return []Message{
		&IDName{Name: a.e.Name()},
		&IDAuthor{Author: a.e.Author()},
		&UCIOk{},
	}, nil
}
