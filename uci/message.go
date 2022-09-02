package uci

import (
	"bytes"
	"encoding"
	"fmt"
	"strings"
)

type ErrUnknownMessage struct {
	Message Message
}

func (e ErrUnknownMessage) Error() string {
	return fmt.Sprintf("unknown message: %s", e.Message)
}

type ErrUnknownCommand struct {
	Command string
}

func (e ErrUnknownCommand) Error() string {
	return fmt.Sprintf("unknown command: %s", e.Command)
}

type ErrCannotUnmarshal struct {
	Text []byte
}

func (e ErrCannotUnmarshal) Error() string {
	return fmt.Sprintf("cannot unmarshal: %s", e.Text)
}

type Message interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	EngineBound() bool
}

// Parse tries to parse a line into an engine-bound UCI message.
func Parse(line string) (Message, error) {
	// Ignore empty lines.
	if line == "" {
		return nil, nil
	}

	words := strings.Fields(line)
	command := words[0]

	switch command {
	default:
		return nil, ErrUnknownCommand{command}
	case "uci":
		var m UCI
		err := m.UnmarshalText([]byte(line))
		return &m, err
	case "id":
		if len(words) == 1 {
			return nil, ErrUnknownCommand{command}
		}
		switch words[1] {
		case "name":
			var m IDName
			err := m.UnmarshalText([]byte(line))
			return &m, err
		case "author":
			var m IDAuthor
			err := m.UnmarshalText([]byte(line))
			return &m, err
		default:
			return nil, ErrUnknownCommand{command}
		}
	}
}

// UCI represents the "uci" command.
type UCI struct{}

func (u *UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

func (u *UCI) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("uci")) {
		var err ErrCannotUnmarshal
		copy(err.Text, text)
		return err
	}

	return nil
}

func (u *UCI) EngineBound() bool {
	return true
}

// IDName represents the "id name" command.
type IDName struct {
	Name string
}

func (i *IDName) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("id name %s", i.Name)), nil
}

func (i *IDName) UnmarshalText(text []byte) error {
	if !bytes.HasPrefix(text, []byte("id name ")) {
		var err ErrCannotUnmarshal
		copy(err.Text, text)
		return err
	}

	i.Name = strings.TrimPrefix(string(text), "id name ")
	return nil
}

func (i *IDName) EngineBound() bool {
	return true
}

// IDAuthor represents the "id author" command.
type IDAuthor struct {
	Author string
}

func (i *IDAuthor) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("id author %s", i.Author)), nil
}

func (i *IDAuthor) UnmarshalText(text []byte) error {
	if !bytes.HasPrefix(text, []byte("id author ")) {
		var err ErrCannotUnmarshal
		copy(err.Text, text)
		return err
	}

	i.Author = strings.TrimPrefix(string(text), "id author ")
	return nil
}

func (i *IDAuthor) EngineBound() bool {
	return true
}

// UCIOk represents the "uciok" command.
type UCIOk struct{}

func (u *UCIOk) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

func (u *UCIOk) UnmarshalText(text []byte) error {
	if !bytes.Equal(text, []byte("uciok")) {
		var err ErrCannotUnmarshal
		copy(err.Text, text)
		return err
	}

	return nil
}

func (u *UCIOk) EngineBound() bool {
	return true
}
