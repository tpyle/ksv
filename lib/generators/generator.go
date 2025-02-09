package generators

type Generator interface {
	Generate(params map[string]string) (string, error)
	GetRef() string
}

func GetGenerator(ref string) (bool, Generator) {
	switch ref {
	case "owasp":
		return true, &OWASPGenerator{}
	case "number":
		return true, &NumberGenerator{}
	case "uuid":
		return true, &UUIDGenerator{}
	default:
		return false, &UnsupportedGenerator{}
	}
}
