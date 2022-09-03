package chess

// CastlingRights is a bitset of available castling rights.
type CastlingRights uint8

// Castling rights constants.
const (
	WhiteOO CastlingRights = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

func NewCastlingRights() CastlingRights {
	return WhiteOO | WhiteOOO | BlackOO | BlackOOO
}

func (c *CastlingRights) Contains(other CastlingRights) bool {
	return *c&other != 0
}

func (c *CastlingRights) Remove(other CastlingRights) {
	*c &^= other
}

func (c *CastlingRights) Add(other CastlingRights) {
	*c |= other
}
