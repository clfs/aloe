package uci

import (
	"encoding"
	"fmt"
	"strings"
)

// Message is the interface implemented by all UCI messages.
type Message interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// BestMove represents the "bestmove" message.
type BestMove struct {
	Move   string
	Ponder string
}

func (b *BestMove) MarshalText() ([]byte, error) {
	var text []byte

	if b.Move == "" {
		return nil, fmt.Errorf("invalid bestmove")
	}

	text = fmt.Appendf(text, "bestmove %s", b.Move)

	if b.Ponder != "" {
		text = fmt.Appendf(text, " ponder %s", b.Ponder)
	}

	return text, nil
}

func (b *BestMove) UnmarshalText(text []byte) error {
	fields := strings.Fields(string(text))

	if len(fields) == 0 {
		return fmt.Errorf("invalid bestmove")
	}

	if fields[0] != "bestmove" {
		return fmt.Errorf("invalid bestmove")
	}

	if len(fields) < 2 {
		return fmt.Errorf("invalid bestmove")
	}

	b.Move = fields[1]

	if len(fields) > 2 {
		if fields[2] != "ponder" {
			return fmt.Errorf("invalid bestmove")
		}

		if len(fields) < 4 {
			return fmt.Errorf("invalid bestmove")
		}

		b.Ponder = fields[3]
	}

	return nil
}

// CopyProtection represents the "copyprotection" message.
type CopyProtection struct{}

func (c *CopyProtection) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *CopyProtection) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Debug represents the "debug" message.
type Debug struct {
	Flag bool
}

func (d *Debug) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Debug) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// Go represents the "go" message.
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

func (g *Go) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// ID represents the "id" message.
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

// Info represents the "info" message.
type Info struct{}

func (i *Info) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (i *Info) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// IsReady represents the "isready" message.
type IsReady struct{}

func (*IsReady) MarshalText() ([]byte, error) {
	return []byte("isready"), nil
}

func (*IsReady) UnmarshalText(text []byte) error {
	if string(text) != "isready" {
		return fmt.Errorf("invalid isready message")
	}
	return nil
}

// Option represents the "option" message.
type Option struct{}

func (o *Option) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (o *Option) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// PonderHit represents the "ponderhit" message.
type PonderHit struct{}

func (*PonderHit) MarshalText() ([]byte, error) {
	return []byte("ponderhit"), nil
}

func (*PonderHit) UnmarshalText(text []byte) error {
	if string(text) != "ponderhit" {
		return fmt.Errorf("invalid ponderhit message")
	}
	return nil
}

// Position represents the "position" message.
type Position struct {
	FEN   string
	Moves []string
}

func (p *Position) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *Position) UnmarshalText(text []byte) error {
	return fmt.Errorf("not message")
}

// Quit represents the "quit" message.
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

// ReadyOk represents the "readyok" message.
type ReadyOk struct{}

func (*ReadyOk) MarshalText() ([]byte, error) {
	return []byte("readyok"), nil
}

func (*ReadyOk) UnmarshalText(text []byte) error {
	if string(text) != "readyok" {
		return fmt.Errorf("invalid readyok message")
	}
	return nil
}

// Register represents the "register" message.
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

// Registration represents the "registration" message.
type Registration struct{}

func (r *Registration) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Registration) UnmarshalText(text []byte) error {
	return fmt.Errorf("not implemented")
}

// SetOption represents the "setoption" message.
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

// Stop represents the "stop" message.
type Stop struct{}

func (*Stop) MarshalText() ([]byte, error) {
	return []byte("stop"), nil
}

func (*Stop) UnmarshalText(text []byte) error {
	if string(text) != "stop" {
		return fmt.Errorf("invalid stop message")
	}
	return nil
}

// UCI represents the "uci" message.
type UCI struct{}

func (*UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

func (*UCI) UnmarshalText(text []byte) error {
	if string(text) != "uci" {
		return fmt.Errorf("invalid uci message")
	}
	return nil
}

// UCINewGame represents the "ucinewgame" message.
type UCINewGame struct{}

func (*UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

func (*UCINewGame) UnmarshalText(text []byte) error {
	if string(text) != "ucinewgame" {
		return fmt.Errorf("invalid ucinewgame message")
	}
	return nil
}

// UCIOk represents the "uciok" message.
type UCIOk struct{}

func (*UCIOk) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

func (*UCIOk) UnmarshalText(text []byte) error {
	if string(text) != "uciok" {
		return fmt.Errorf("invalid uciok message")
	}
	return nil
}

// Unknown represents a message of unknown type.
type Unknown struct {
	Text string
}

func (u *Unknown) MarshalText() ([]byte, error) {
	return []byte(u.Text), nil
}

func (u *Unknown) UnmarshalText(text []byte) error {
	u.Text = string(text)
	return nil
}
