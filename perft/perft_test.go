package perft

import (
	"testing"

	"github.com/clfs/aloe/chess"
)

func BenchmarkCount_5(b *testing.B) {
	p := chess.NewPosition()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Count(&p, 5)
	}
}

func BenchmarkDivide_5(b *testing.B) {
	p := chess.NewPosition()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Divide(&p, 5)
	}
}
