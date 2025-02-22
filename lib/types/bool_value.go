package types

import (
	"strconv"

	"github.com/tpyle/ksv/lib/errors"
)

type KSVBool bool

func (b *KSVBool) Get(path string) (Queryable, error) {
	if path != "" {
		return nil, errors.ErrInvalidPath
	}
	return b, nil
}

func (b *KSVBool) Set(value string) error {
	bval, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*b = KSVBool(bval)
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
	return strconv.FormatBool(bool(*b))
}
