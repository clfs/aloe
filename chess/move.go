package chess

import "fmt"

// Move is a chess move. The zero value for Move is a null move.
type Move struct {
	// The starting square. When castling, this is the king's starting square.
	From Square

	// The destination square. When castling, this is the king's destination square.
	To Square

	// Promotion information. For non-promotion moves, this is [NoPromotion].
	PromotionInfo PromotionInfo
}

// NewMove accepts UCI-compatible long algebraic notation (LAN) and returns the
// corresponding move.
func NewMove(s string) (Move, error) {
	if len(s) < 4 || len(s) > 5 {
		return Move{}, fmt.Errorf("invalid move")
	}

	from, err := parseSquare(s[:2])
	if err != nil {
		return Move{}, fmt.Errorf("invalid move: %v", err)
	}

	to, err := parseSquare(s[2:4])
	if err != nil {
		return Move{}, fmt.Errorf("invalid move: %v", err)
	}

	var promo PromotionInfo

	if len(s) == 5 {
		switch ch := s[4]; ch {
		case 'q':
			promo = QueenPromotion
		case 'r':
			promo = RookPromotion
		case 'b':
			promo = BishopPromotion
		case 'n':
			promo = KnightPromotion
		default:
			return Move{}, fmt.Errorf("invalid move: invalid promotion %c", ch)
		}
	}

	return Move{from, to, promo}, nil
}
