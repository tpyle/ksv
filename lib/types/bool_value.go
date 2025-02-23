package types

import (
	"strconv"

	"github.com/tpyle/ksv/lib/ksverrors"
)

var (
	KSVBoolTrue  = KSVBool{Value: true}
	KSVBoolFalse = KSVBool{Value: false}
)

type KSVBool struct {
	Value bool
}

func (b *KSVBool) Get(path string) (Queryable, error) {
	if path != "" {
		return nil, ksverrors.ErrInvalidPath
	}
	return b, nil
}

func (b *KSVBool) Set(value string) error {
	bval, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	b.Value = bval
	return nil
}

func (b *KSVBool) Validate() []error {
	return nil
}

func (b *KSVBool) GetChildren() []string {
	return []string{""}
}

func (b *KSVBool) GetValues() map[string]string {
	return map[string]string{"": b.String()}
}

func (b *KSVBool) String() string {
	return strconv.FormatBool(b.Value)
}

func (b *KSVBool) ToBool() bool {
	return b.Value
}

func NewKSVBool(value bool) *KSVBool {
	return &KSVBool{Value: value}
}
