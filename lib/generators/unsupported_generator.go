package generators

import (
	"github.com/tpyle/ksv/lib/ksverrors"
)

type UnsupportedGenerator struct {
}

func (n *UnsupportedGenerator) Generate(params map[string]string) (string, error) {
	return "", ksverrors.ErrUnsupportedGenerator
}

func (n *UnsupportedGenerator) GetRef() string {
	return "unsupported"
}

func (n *UnsupportedGenerator) GetDefaultParams() map[string]string {
	return map[string]string{}
}
