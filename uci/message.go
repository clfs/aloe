package uci

import (
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

type Message interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	EngineBound() bool
}

// parse tries to parse a line into an engine-bound UCI message.
func parse(line string) (Message, error) {
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
		return &UCI{}, nil
	}
}

// UCI represents the "uci" command.
type UCI struct{}

func (u *UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

func (u *UCI) UnmarshalText(_ []byte) error {
	return nil
}

func (u *UCI) EngineBound() bool {
	return true
}

// ID represents the "id" command.
type ID struct {
	Name   string
	Author string
}

func (i *ID) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("id name %s\nid author %s", i.Name, i.Author)), nil
}

func (i *ID) UnmarshalText(text []byte) error {
	return nil // TODO
}

func (i *ID) EngineBound() bool {
	return false
}

// UCIOk represents the "uciok" command.
type UCIOk struct{}

func (u *UCIOk) MarshalText() ([]byte, error) {
	return []byte("uciok"), nil
}

func (u *UCIOk) UnmarshalText(_ []byte) error {
	return nil
}

func (u *UCIOk) EngineBound() bool {
	return false
}
