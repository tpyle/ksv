package generators

import (
	"github.com/tpyle/ksv/lib/errors"
)

type UnsupportedGenerator struct {
}

func (n *UnsupportedGenerator) Generate(params map[string]string) (string, error) {
	return "", errors.ErrUnsupportedGenerator
}

func (n *UnsupportedGenerator) GetRef() string {
	return "unsupported"
}
