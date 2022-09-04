package chess

// PromotionInfo represents promotion information for a move.
type PromotionInfo uint8

// [PromotionInfo] constants.
const (
	NoPromotion PromotionInfo = iota
	KnightPromotion
	BishopPromotion
	RookPromotion
	QueenPromotion
)

// Role returns the role that corresponds to the promotion information.
// For [NoPromotion], ok is false.
func (p PromotionInfo) Role() (r Role, ok bool) {
	if p == NoPromotion {
		return 0, false
	}
	return Role(p), true
}
