package generators

type Generator interface {
	Generate(params map[string]string) (string, error)
}
