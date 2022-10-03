package uci

import (
	"encoding"
	"fmt"
	"io"
)

// Encoder encodes UCI messages to an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes m to the stream, followed by a newline character.
func (e *Encoder) Encode(m encoding.TextMarshaler) error {
	text, err := m.MarshalText()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(e.w, "%s\n", text)
	return err
}
