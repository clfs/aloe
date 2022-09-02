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
