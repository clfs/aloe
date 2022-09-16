// Package uci implements the Universal Chess Interface (UCI) protocol.
package uci

import "context"

// Engine is the interface that a chess engine must implement for compatibility
// with this package. Aloe's engine implements this interface.
type Engine interface {
	UCIID() ID
	UCIGo(ctx context.Context, g Go) (Info, error)
}

// ID represents the "id" UCI command.
type ID struct {
	Name    string
	Authors []string
}

// Go represents the "go" UCI command.
type Go struct {
	FEN        string
	Parameters Parameters

	// Engines may implement streaming by running this function on
	// each intermediate search result.
	OnUpdate func(Info)
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
	PV        []string // The best line found ("principal variation").
	Score     int      // The score from the engine's point of view.
	ScoreType string   // Either ScoreTypeCentipawn or ScoreTypeMate.
}
