package chess

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
