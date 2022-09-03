package chess

// Promotion represents promotion information for a move.
type Promotion uint8

// Promotion constants.
const (
	NoPromotion Promotion = iota
	KnightPromotion
	BishopPromotion
	RookPromotion
	QueenPromotion
)

// Role returns the role that corresponds to the promotion information.
// For NoPromotion, ok is false.
func (p Promotion) Role() (r Role, ok bool) {
	if p == NoPromotion {
		return 0, false
	}
	return Role(p), true
}
