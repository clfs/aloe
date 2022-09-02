package engine

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Name() string {
	return "aloe"
}

func (e *Engine) Author() string {
	return "Calvin Figuereo-Supraner"
}

func (e *Engine) NewGame() error {
	return nil // TODO: implement.
}

func (e *Engine) NewGameFromFEN(fen string) error {
	return nil // TODO: implement.
}

func (e *Engine) MoveAlgebraic(move string) error {
	return nil // TODO: implement.
}
