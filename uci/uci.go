// Package uci implements the Universal Chess Interface (UCI) protocol.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/clfs/aloe/fen"
)

// Engine is the interface that a chess engine must implement for compatibility
// with this package. Aloe's engine implements this interface.
type Engine interface {
	ID() ID
	Search(p Position, g Go, ch <-chan Info) error
}

// ID represents the "id" UCI command.
type ID struct {
	Name   string
	Author string
}

// Position represents the "position" UCI command.
type Position struct {
	FEN   string
	Moves []string
}

// DefaultPosition is the default [Position] used when no "position" command is
// available.
var DefaultPosition = Position{FEN: fen.StartingFEN}

// Regexps for parsing UCI "position" commands.
var (
	rgxPositionStartposMoves = regexp.MustCompile(`^position startpos moves (.+)$`)
	rgxPositionFENMoves      = regexp.MustCompile(`^position fen (.+) moves (.+)$`)
	rgxPositionFEN           = regexp.MustCompile(`^position fen (.+)$`)
)

func (p *Position) UnmarshalText(text []byte) error {
	s := string(text)

	if s == "position startpos" {
		*p = Position{fen.StartingFEN, nil}
		return nil
	}

	// position startpos moves <moves>
	if m := rgxPositionStartposMoves.FindStringSubmatch(s); m != nil {
		*p = Position{fen.StartingFEN, strings.Fields(m[1])}
		return nil
	}

	// position fen <fen> moves <moves>
	if m := rgxPositionFENMoves.FindStringSubmatch(s); m != nil {
		*p = Position{m[1], strings.Fields(m[2])}
		return nil
	}

	// position fen <fen>
	if m := rgxPositionFEN.FindStringSubmatch(s); m != nil {
		*p = Position{m[1], nil}
		return nil
	}

	return fmt.Errorf("invalid position command: %s", text)
}

// Go represents the "go" UCI command.
type Go struct {
	SearchMoves []string // Restrict search to these moves only. Ignore if empty.

	Ponder   bool // Search in pondering mode.
	Infinite bool // Search until interrupted.
	MoveTime int  // If > 0, search for this many milliseconds.

	WhiteTime      int // If > 0, white's remaining time in milliseconds.
	BlackTime      int // If > 0, black's remaining time in milliseconds.
	WhiteIncrement int // If > 0, white's increment in milliseconds.
	BlackIncrement int // If > 0, black's increment in milliseconds.

	Depth     int // If > 0, search this many plies only.
	Nodes     int // If > 0, search this many nodes only.
	Mate      int // If > 0, search for a mate in this many moves.
	MovesToGo int // If > 0, there are this many moves until the next time control.
}

func (g *Go) UnmarshalText(text []byte) error {
	return nil // TODO: implement
}

// Score types used in [Info].
const (
	ScoreTypeCentipawn = "cp"
	ScoreTypeMate      = "mate"
)

// Info represents the "info" UCI command.
type Info struct {
	Depth     int      // Search depth in plies.
	PV        []string // Best line found ("principal variation").
	Score     int      // Score from the engine's point of view.
	ScoreType string   // Either ScoreTypeCentipawn or ScoreTypeMate.
}

func (i *Info) String() string {
	return "todo" // implement
}

// Client is a wrapper around a UCI-compatible engine.
type Client struct {
	e  Engine
	w  io.Writer // For UCI output.
	ch chan Info // Channel for search info.
	p  Position  // Position command to use for search.
}

// NewClient returns a new [Client] that writes UCI responses to w.
func NewClient(e Engine, w io.Writer) *Client {
	return &Client{
		e:  e,
		w:  w,
		ch: make(chan Info),
		p:  DefaultPosition,
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
	c.p = DefaultPosition
}

// handlePosition handles the "position" UCI command.
func (c *Client) handlePosition(line string) error {
	return c.p.UnmarshalText([]byte(line))
}

// handleGo handles the "go" UCI command.
func (c *Client) handleGo(line string) error {
	close(c.ch) // Cancel the existing search, if any.
	c.ch = make(chan Info)

	var g Go
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
			fmt.Fprintln(c.w, info)
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
