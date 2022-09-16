// Package uci implements the Universal Chess Interface (UCI) protocol.
package uci

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/clfs/aloe/chess"
	"github.com/clfs/aloe/fen"
)

// Engine is the interface that a chess engine must implement for compatibility
// with this package. Aloe's engine implements this interface.
type Engine interface {
	UCIID() ID
	UCIGo(g Go, ch <-chan Info) error
}

// ID represents the "id" UCI command.
type ID struct {
	Name   string
	Author string
}

// Go represents the "go" UCI command. It's also tagged with the FEN to analyze.
type Go struct {
	FEN        string
	Parameters Parameters
}

// Parameters describes search parameters for the "go" UCI command.
type Parameters struct {
	Moves []string // Restrict search to these moves only. Ignore if empty.

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

// Client is a wrapper around a UCI-compatible engine.
type Client struct {
	e   Engine
	w   io.Writer // For UCI output.
	ch  chan Info // Channel for search info.
	fen string    // Position under analysis.
}

// NewClient returns a new [Client] that writes UCI responses to w.
func NewClient(e Engine, w io.Writer) *Client {
	return &Client{
		e:   e,
		w:   w,
		fen: fen.StartingFEN,
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
	id := c.e.UCIID()
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
	if c.ch != nil {
		close(c.ch)
	}
	c.fen = fen.StartingFEN
}

// handlePosition handles the "position" UCI command.
func (c *Client) handlePosition(line string) error {
	fen, err := parsePosition(line)
	if err != nil {
		return err
	}
	c.fen = fen
	return nil
}

// handleGo handles the "go" UCI command.
func (c *Client) handleGo(line string) error {
	if c.ch != nil {
		return fmt.Errorf("search already in progress")
	}
	return nil // TODO: implement
}

// handleStop handles the "stop" UCI command.
func (c *Client) handleStop() {
	if c.ch != nil {
		close(c.ch)
	}
}

// handleQuit handles the "quit" UCI command.
func (c *Client) handleQuit() {
	if c.ch != nil {
		close(c.ch)
	}
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

// parsePosition parses a "position" UCI command and returns the FEN string it represents.
func parsePosition(line string) (string, error) {
	var res string // The FEN to return.

	fields := strings.Fields(line)

	// A valid position command must be at least 2 words long, and the first
	// word must be "position".
	if len(fields) < 2 || fields[0] != "position" {
		return "", fmt.Errorf("invalid position command: %s", line)
	}
	fields = fields[1:]

	// The next word must be "startpos" or "fen". If it's "fen", read farther
	// ahead to get the initial FEN.
	switch fields[0] {
	case "startpos":
		res = fen.StartingFEN
		fields = fields[2:]
	case "fen":
		res = strings.Join(fields[2:8], " ")
		fields = fields[8:]
	default:
		return "", fmt.Errorf("invalid position command: %s", line)
	}

	// If there are remaining words, they must be moves to apply to the initial
	// FEN. Notably, providing "moves" as the final word is interpreted as an
	// empty move list.
	var moves []string

	if len(fields) > 0 {
		if fields[0] != "moves" {
			return "", fmt.Errorf("invalid position command: %s", line)
		}
		if len(fields) > 1 {
			moves = fields[1:]
		}
	}

	// Decode the current FEN.
	pos, err := fen.Decode(res)
	if err != nil {
		return "", fmt.Errorf("invalid position command: %v", err)
	}

	// Apply moves, if any.
	for _, m := range moves {
		move, err := chess.NewMove(m)
		if err != nil {
			return "", fmt.Errorf("invalid position command: %v", err)
		}
		// TODO: This isn't guaranteed to be legal.
		_ = pos.Move(move)
	}

	// Encode the resulting position back to a FEN and return it.
	res, err = fen.Encode(pos)
	if err != nil {
		return "", fmt.Errorf("invalid position command: %v", err)
	}
	return res, nil
}
