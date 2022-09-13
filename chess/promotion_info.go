package chess

import "fmt"

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

// IsValid returns true if p is a package-defined PromotionInfo constant.
func (p PromotionInfo) IsValid() bool {
	return p <= QueenPromotion
}

func (p PromotionInfo) String() string {
	switch p {
	case NoPromotion:
		return "NoPromotion"
	case KnightPromotion:
		return "KnightPromotion"
	case BishopPromotion:
		return "BishopPromotion"
	case RookPromotion:
		return "RookPromotion"
	case QueenPromotion:
		return "QueenPromotion"
	default:
		return fmt.Sprintf("PromotionInfo(%d)", p)
	}
}
