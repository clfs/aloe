package uci

import (
	"bytes"
	"testing"
)

func FuzzDecoderMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		dec := NewDecoder(bytes.NewReader(b))
		for {
			msg, err := dec.Message()
			if err != nil {
				break
			}
			_ = msg
		}
	})
}
