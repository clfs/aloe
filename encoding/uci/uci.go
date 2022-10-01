// Package uci implements encoding and decoding for Universal Chess Interface
// (UCI) messages.
package uci

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Message holds a value of one of these types:
//   - [BestMove]
//   - [CopyProtection]
//   - [Debug]
//   - [Go]
//   - [ID]
//   - [Info]
//   - [IsReady]
//   - [Option]
//   - [PonderHit]
//   - [Position]
//   - [Quit]
//   - [ReadyOk]
//   - [Register]
//   - [Registration]
//   - [SetOption]
//   - [Stop]
//   - [UCI]
//   - [UCINewGame]
//   - [Unknown]
type Message any

// Marshaler is the interface implemented by types that can marshal themselves
// into a valid UCI message.
type Marshaler interface {
	MarshalUCI() ([]byte, error)
}

// A MarshalerError represents an error from calling a MarshalUCI method.
type MarshalerError struct {
	Type reflect.Type
	Err  error
}

func (e *MarshalerError) Error() string {
	return fmt.Sprintf("error marshaling for type %s: %s", e.Type, e.Err)
}

// Unmarshaler is the interface implemented by types that can unmarshal a UCI
// description of themselves. The input can be assumed to be a valid encoding of
// a UCI value. UnmarshalUCI must copy the UCI data if it wishes to retain the
// data after returning.
type Unmarshaler interface {
	UnmarshalUCI([]byte) error
}

// An UnmarshalerError represents an error from calling an UnmarshalUCI method.
type UnmarshalerError struct {
	Type reflect.Type
	Err  error
}

func (e *UnmarshalerError) Error() string {
	return fmt.Sprintf("error unmarshaling for type %s: %s", e.Type, e.Err)
}

// Decoder decodes UCI messages from an input stream.
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r beyond the
// values requested.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode reads the next UCI-encoded value from its input and stores it in the
// value pointed to by m.
func (dec *Decoder) Decode(v Unmarshaler) error {
	return nil
}

// Message returns the next UCI message in the input stream. At the end of the
// input stream, Message returns nil, io.EOF.
func (dec *Decoder) Message() (Message, error) {
	return nil, io.EOF
}

// Encoder encodes UCI messages to an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes the UCI encoding of m to the stream, followed by a newline
// character.
func (enc *Encoder) Encode(v Marshaler) error {
	text, err := v.MarshalUCI()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(enc.w, "%s\n", text)
	return err
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
	Text []byte
}

func (u *Unknown) MarshalText() ([]byte, error) {
	return u.Text, nil
}

func (u *Unknown) UnmarshalText(text []byte) error {
	u.Text = nil
	copy(u.Text, text)
	return nil
}
