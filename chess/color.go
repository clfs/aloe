package chess

// Color is the color of a player, piece, square, or other item. A Color is
// either White or Black.
type Color int8

// Color constants.
const (
	White Color = iota
	Black
)
