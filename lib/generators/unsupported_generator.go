package generators

import (
	"github.com/tpyle/ksv/lib/ksverrors"
)

type UnsupportedGenerator struct {
	Ref string
}

func (n *UnsupportedGenerator) Generate(params map[string]string) (string, error) {
	return "", ksverrors.ErrUnsupportedGenerator
}

func (n *UnsupportedGenerator) GetRef() string {
	return n.Ref
}

func (n *UnsupportedGenerator) GetDefaultParams() map[string]string {
	return map[string]string{}
}
