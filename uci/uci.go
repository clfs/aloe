package uci

import (
	"bufio"
	"io"

	"github.com/clfs/aloe/engine"
)

type Adapter struct {
	e *engine.Engine
	r io.Reader
	w io.Writer
}

func NewAdapter(e *engine.Engine, r io.Reader, w io.Writer) *Adapter {
	return &Adapter{e, r, w}
}

func (a *Adapter) Run() error {
	scanner := bufio.NewScanner(a.r)

	for scanner.Scan() {
		line := scanner.Text()

		request, err := parse(line)
		if err != nil {
			return err
		}

		responses, err := a.handle(request)
		if err != nil {
			return err
		}

		for _, r := range responses {
			cmd, err := r.MarshalText()
			if err != nil {
				return err
			}

			if _, err := a.w.Write(cmd); err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func (a *Adapter) handle(req Request) ([]Response, error) {
	switch v := req.(type) {
	case *UCI:
		return a.handleUCI(v)
	}

	return nil, ErrInvalidRequest{[]byte("todo")}
}

func (a *Adapter) handleUCI(req *UCI) ([]Response, error) {
	return []Response{&ID{}}, nil
}
