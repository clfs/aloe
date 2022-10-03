package uci

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Decoder decodes UCI messages from an input stream.
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: bufio.NewReader(r)}
}

// Decode reads the next message from its input and stores it in the value
// pointed to by m.
func (d *Decoder) Decode(m Message) error {
	text, err := d.r.ReadBytes('\n')
	if err != nil {
		return err
	}
	return m.UnmarshalText(text[:len(text)-1])
}

// Message returns the next UCI message in the input stream. At the end of the
// input stream, Message returns nil, io.EOF.
func (d *Decoder) Message() (Message, error) {
	text, err := d.r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return nil, fmt.Errorf("blank message")
	}

	return nil, io.EOF
}
