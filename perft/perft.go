package perft

import "github.com/clfs/aloe/chess"

// Count walks the legal move tree for a position and returns the number of
// leaf nodes at the given depth.
//
// Count does not check for:
//   - draws by insufficient material
//   - draws by three-fold repetition
//   - draws by 50-move rule
//
// If depth is non-positive, Count always returns 1.
func Count(p *chess.Position, depth int) int {
	if depth < 1 {
		return 1
	}

	moves := p.LegalMoves()

	if depth == 1 {
		return len(moves)
	}

	var nodes int

	for _, m := range moves {
		undo := p.Move(m)
		nodes += Count(p, depth-1)
		p.Undo(undo)
	}

	return nodes
}

// Divide computes [Count] for each possible position after the provided one,
// with 1 fewer depth. The result is keyed by the move that led to the position.
func Divide(p *chess.Position, depth int) map[chess.Move]int {
	moves := p.LegalMoves()

	nodes := make(map[chess.Move]int)

	for _, m := range moves {
		undo := p.Move(m)
		nodes[m] = Count(p, depth-1)
		p.Undo(undo)
	}

	return nodes
}
