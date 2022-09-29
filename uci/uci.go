// Package uci implements the Universal Chess Interface (UCI) protocol.
package uci

import (
	"bufio"
	"io"
)

type Engine interface {
}

// Client is a wrapper around a UCI-compatible engine.
type Client struct {
	e Engine
	w io.Writer
}

// NewClient returns a new [Client] that writes UCI output to w.
func NewClient(e Engine, w io.Writer) *Client {
	return &Client{e: e, w: w}
}

// Run parses UCI input from r and sends commands to the engine.
func (c *Client) Run(r io.Reader) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()

		req, err := Parse(line)
		if err != nil {
			return err
		}

		_ = req // discard for now, TODO
	}

	return s.Err()
}
