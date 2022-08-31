package uci

import (
	"encoding"
	"fmt"
)

type Response interface {
	encoding.TextMarshaler
}

type ID struct {
	Name   string
	Author string
}

func (i *ID) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("id name %s\nid author %s\n", i.Name, i.Author)), nil
}
