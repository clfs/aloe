// Package uci implements encoding and decoding for Universal Chess Interface
// (UCI) messages.
package uci

import "fmt"

// BestMove represents the "bestmove" command.
type BestMove struct {
	Move   string
	Ponder string
}

func (b *BestMove) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (b *BestMove) UnmarshalText() error {
	return fmt.Errorf("not implemented")
}

// CopyProtection represents the "copyprotection" command.
type CopyProtection struct{}

func (c *CopyProtection) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *CopyProtection) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Debug represents the "debug" command.
type Debug struct {
	Flag bool
}

func (d *Debug) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Debug) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Go represents the "go" command.
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

func (g *Go) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

// ID represents the "id" command.
type ID struct {
	Name   string
	Author string
}

func (i *ID) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *ID) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Info represents the "info" UCI command.
type Info struct{}

func (i *Info) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *Info) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// IsReady represents the "isready" command.
type IsReady struct{}

func (*IsReady) MarshalText() ([]byte, error) {
	return []byte("isready"), nil
}

func (*IsReady) UnmarshalText(text []byte) error {
	if string(text) != "isready" {
		return fmt.Errorf("invalid isready command")
	}
	return nil
}

// Option represents the "option" command.
type Option struct{}

func (o *Option) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (o *Option) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// PonderHit represents the "ponderhit" command.
type PonderHit struct{}

func (*PonderHit) MarshalText() ([]byte, error) {
	return []byte("ponderhit"), nil
}

func (*PonderHit) UnmarshalText(text []byte) error {
	if string(text) != "ponderhit" {
		return fmt.Errorf("invalid ponderhit command")
	}
	return nil
}

// Position represents the "position" command.
type Position struct {
	FEN   string
	Moves []string
}

func (p *Position) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *Position) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Quit represents the "quit" command.
type Quit struct{}

func (*Quit) MarshalText() ([]byte, error) {
	return []byte("quit"), nil
}

func (*Quit) UnmarshalText(text []byte) error {
	if string(text) != "quit" {
		return fmt.Errorf("invalid quit command")
	}
	return nil
}

// ReadyOk represents the "readyok" command.
type ReadyOk struct{}

func (*ReadyOk) MarshalText() ([]byte, error) {
	return []byte("readyok"), nil
}

func (*ReadyOk) UnmarshalText(text []byte) error {
	if string(text) != "readyok" {
		return fmt.Errorf("invalid readyok command")
	}
	return nil
}

// Register represents the "register" command.
type Register struct {
	Later bool
	Name  string
	Code  string
}

func (r *Register) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Register) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Registration represents the "registration" command.
type Registration struct{}

func (r *Registration) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Registration) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// SetOption represents the "setoption" command.
type SetOption struct {
	Name  string
	Value string
}

func (s *SetOption) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *SetOption) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Stop represents the "stop" command.
type Stop struct{}

func (*Stop) MarshalText() ([]byte, error) {
	return []byte("stop"), nil
}

func (*Stop) UnmarshalText(text []byte) error {
	if string(text) != "stop" {
		return fmt.Errorf("invalid stop command")
	}
	return nil
}

// UCI represents the "uci" command.
type UCI struct{}

func (*UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

func (*UCI) UnmarshalText(text []byte) error {
	if string(text) != "uci" {
		return fmt.Errorf("invalid uci command")
	}
	return nil
}

// UCINewGame represents the "ucinewgame" command.
type UCINewGame struct{}

func (*UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

func (*UCINewGame) UnmarshalText(text []byte) error {
	if string(text) != "ucinewgame" {
		return fmt.Errorf("invalid ucinewgame command")
	}
	return nil
}

// UCIOk represents the "uciok" command.
type UCIOk struct{}

func (*UCIOk) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

func (*UCIOk) UnmarshalText(text []byte) error {
	if string(text) != "uciok" {
		return fmt.Errorf("invalid uciok command")
	}
	return nil
}
