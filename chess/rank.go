package chess

import "fmt"

// Rank is a rank on a chess board. Note that [Rank1] = 0.
type Rank uint8

// Board ranks.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) IsValid() bool {
	return r <= Rank8
}

func (r Rank) String() string {
	if !r.IsValid() {
		return fmt.Sprintf("Rank(%d)", r)
	}

	return fmt.Sprintf("Rank%d", r+1)
}
