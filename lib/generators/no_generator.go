package generators

import "github.com/tpyle/ksv/lib/ksverrors"

const (
	NoGeneratorRef = "no_generator"
)

type NoGenerator struct {
}

func (n *NoGenerator) Generate(params map[string]string) (string, error) {
	return "", ksverrors.ErrNoGeneratorCannotGenerate
}

func (n *NoGenerator) GetRef() string {
	return NoGeneratorRef
}

func (n *NoGenerator) GetDefaultParams() map[string]string {
	return map[string]string{}
}
