package composition

type Engine interface {
	getEngineName() string
}

type StandardEngine struct{}
type SportEngine struct{}

func (SportEngine) getEngineName() string {
	return "SportEngine"
}

func (StandardEngine) getEngineName() string {
	return "StandardEngine"
}
