package perft

import (
	"fmt"
	"testing"

	"github.com/clfs/aloe/chess"
)

func decode(t *testing.T, fen string) *chess.Position {
	t.Helper()
	var pos chess.Position
	if err := pos.UnmarshalText([]byte(fen)); err != nil {
		t.Fatal(err)
	}
	return &pos
}

var countTestCases = []struct {
	fen   string
	nodes []int
}{
	{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		[]int{1, 20, 400, 8902, 197281, 4865609, 119060324},
	},
	{
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -",
		[]int{1, 48, 2039, 97862, 4085603, 193690690},
	},
	{
		"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - ",
		[]int{1, 14, 191, 2812, 43238, 674624, 11030083, 178633661},
	},
	{
		"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
		[]int{1, 6, 264, 9467, 422333, 15833292},
	},
	{
		"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
		[]int{1, 44, 1486, 62379, 2103487, 89941194},
	},
	{
		"r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
		[]int{1, 46, 2079, 89890, 3894594, 164075551},
	},
}

func TestCount(t *testing.T) {
	for _, tc := range countTestCases {
		p := decode(t, tc.fen)

		for depth, want := range tc.nodes {
			got := Count(p, depth)

			if want != got {
				t.Errorf("%s, depth %d: want %d, got %d", tc.fen, depth, want, got)
			}
		}
	}
}

func BenchmarkCount_5(b *testing.B) {
	p := chess.NewPosition()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Count(p, 5)
	}
}

func ExampleCount() {
	p := chess.NewPosition()
	nodes := Count(p, 5)
	fmt.Println(nodes)
	// Output: 4865609
}

var divideTestCases = []struct {
	fen   string
	depth int
	nodes map[chess.Move]int
}{
	{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		5,
		map[chess.Move]int{
			{From: chess.A2, To: chess.A3}: 181046,
			{From: chess.B2, To: chess.B3}: 215255,
			{From: chess.C2, To: chess.C3}: 222861,
			{From: chess.D2, To: chess.D3}: 328511,
			{From: chess.E2, To: chess.E3}: 402988,
			{From: chess.F2, To: chess.F3}: 178889,
			{From: chess.G2, To: chess.G3}: 217210,
			{From: chess.H2, To: chess.H3}: 181044,
			{From: chess.A2, To: chess.A4}: 217832,
			{From: chess.B2, To: chess.B4}: 216145,
			{From: chess.C2, To: chess.C4}: 240082,
			{From: chess.D2, To: chess.D4}: 361790,
			{From: chess.E2, To: chess.E4}: 405385,
			{From: chess.F2, To: chess.F4}: 198473,
			{From: chess.G2, To: chess.G4}: 214048,
			{From: chess.H2, To: chess.H4}: 218829,
			{From: chess.B1, To: chess.A3}: 198572,
			{From: chess.B1, To: chess.C3}: 234656,
			{From: chess.G1, To: chess.F3}: 233491,
			{From: chess.G1, To: chess.H3}: 198502,
		},
	},
}

func TestDivide(t *testing.T) {
	for _, tc := range divideTestCases {
		var pos chess.Position

		err := pos.UnmarshalText([]byte(tc.fen))
		if err != nil {
			t.Fatal(err)
		}

		nodes := Divide(&pos, tc.depth)
		for move, want := range tc.nodes {
			if got, ok := nodes[move]; !ok || got != want {
				t.Errorf("%s, depth %d: want %d, got %d", tc.fen, tc.depth, want, got)
			}
		}
	}
}
