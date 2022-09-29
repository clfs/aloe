// Package uci implements the Universal Chess Interface (UCI) protocol.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Engine is the interface that a chess engine must implement for compatibility
// with this package. Aloe's engine implements this interface.
type Engine interface {
	ID() ResponseID
	Search(p RequestPosition, g RequestGo, ch <-chan ResponseInfo) error
}

// Client is a wrapper around a UCI-compatible engine.
type Client struct {
	e  Engine
	w  io.Writer         // For UCI output.
	ch chan ResponseInfo // Channel for search info.
	p  RequestPosition   // Position command to use for search.
}

// NewClient returns a new [Client] that writes UCI responses to w.
func NewClient(e Engine, w io.Writer) *Client {
	return &Client{
		e:  e,
		w:  w,
		ch: make(chan ResponseInfo),
	}
}

// Run reads commands from r and executes them.
func (c *Client) Run(r io.Reader) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()

		var err error

		switch {
		default:
			err = c.handleUnknown(line)

		// Single-word commands.
		case line == "quit":
			c.handleQuit()
			return nil // early exit!
		case line == "uci":
			c.handleUCI()
		case line == "isready":
			c.handleIsReady()
		case line == "ucinewgame":
			c.handleUCINewGame()
		case line == "stop":
			c.handleStop()

		// Multi-word commands.
		case strings.HasPrefix(line, "position "):
			err = c.handlePosition(line)
		case strings.HasPrefix(line, "go "):
			err = c.handleGo(line)
		}

		if err != nil {
			return err
		}
	}

	return s.Err()
}

// handleUCI handles the "uci" UCI command.
func (c *Client) handleUCI() {
	id := c.e.ID()
	fmt.Fprintf(c.w, "id name %s\n", id.Name)
	fmt.Fprintf(c.w, "id author %s\n", id.Author)
	fmt.Fprintf(c.w, "uciok\n")
}

// handleIsReady handles the "isready" UCI command.
func (c *Client) handleIsReady() {
	fmt.Fprintf(c.w, "readyok\n")
}

// handleUCINewGame handles the "ucinewgame" UCI command.
func (c *Client) handleUCINewGame() {
	close(c.ch)
}

// handlePosition handles the "position" UCI command.
func (c *Client) handlePosition(line string) error {
	return c.p.UnmarshalText([]byte(line))
}

// handleGo handles the "go" UCI command.
func (c *Client) handleGo(line string) error {
	close(c.ch) // Cancel the existing search, if any.
	c.ch = make(chan ResponseInfo)

	var g RequestGo
	if err := g.UnmarshalText([]byte(line)); err != nil {
		return err
	}

	// Send the command to the engine. Once the engine finishes, close the
	// channel to signal that the search is over.
	go func() {
		defer close(c.ch)

		if err := c.e.Search(c.p, g, c.ch); err != nil {
			fmt.Fprintf(c.w, "failed search: %v\n", err)
		}
	}()

	// Write search results to the UCI output as they arrive.
	go func() {
		for info := range c.ch {
			fmt.Fprintf(c.w, "info ")
			fmt.Fprintf(c.w, "depth %d ", info.Depth)
			fmt.Fprintf(c.w, "score %s %d ", info.ScoreType, info.Score)
			fmt.Fprintf(c.w, "pv %s", strings.Join(info.PV, " "))
			fmt.Fprintf(c.w, "\n")
		}
	}()

	return nil
}

// handleStop handles the "stop" UCI command.
func (c *Client) handleStop() {
	close(c.ch)
}

// handleQuit handles the "quit" UCI command.
func (c *Client) handleQuit() {
	close(c.ch)
}

// handleUnknown handles an unknown UCI command.
func (c *Client) handleUnknown(line string) error {
	fields := strings.Fields(line)

	// Ignore empty lines.
	if len(fields) == 0 {
		return nil
	}

	// Ignore valid but unimplemented commands.
	switch fields[0] {
	case "debug", "setoption", "register", "ponderhit":
		return nil
	default:
		return fmt.Errorf("unknown command: %s", fields[0])
	}
}

func Parse(line string) (Request, error) {
	return nil, nil // TODO: implement
}
