package chess

func perft(pos *Position, depth int) int {
	if depth == 0 {
		return 1
	}

	var nodes int
	for _, m := range pos.LegalMoves() {
		undo := pos.Move(m)
		nodes += perft(pos, depth-1)
		pos.Undo(undo)
	}
	return nodes
}
