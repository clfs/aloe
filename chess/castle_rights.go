package chess

// CastleRights is a bitset of available castle rights. The zero value
// indicates no castle rights are available.
type CastleRights uint8

// Castle rights constants.
const (
	WhiteOO CastleRights = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

func NewCastleRights() CastleRights {
	return WhiteOO | WhiteOOO | BlackOO | BlackOOO
}

func (c *CastleRights) Contains(other CastleRights) bool {
	return *c&other != 0
}

func (c *CastleRights) Remove(other CastleRights) {
	*c &^= other
}

func (c *CastleRights) Add(other CastleRights) {
	*c |= other
}
