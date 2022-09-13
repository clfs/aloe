package chess

// Color is the color of a player, piece, square, or other item. A Color is
// either White or Black.
type Color bool

// Color constants.
const (
	White Color = false
	Black Color = true
)

func (c Color) String() string {
	if c {
		return "Black"
	}
	return "White"
}
