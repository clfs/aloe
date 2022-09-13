package chess

import "fmt"

// File is a file on a chess board.
type File uint8

// Board files.
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) IsValid() bool {
	return f <= FileH
}

func (f File) String() string {
	if !f.IsValid() {
		return fmt.Sprintf("File(%d)", f)
	}

	return fmt.Sprintf("File%c", 'A'+f)
}
