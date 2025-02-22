package types

import (
	"github.com/tpyle/ksv/lib/errors"
)

type SecretField struct {
	Value        KSVString          `json:"value"`
	GeneratorRef GeneratorReference `json:"generatorRef"`
}

func (sf *SecretField) Get(path string) (Queryable, error) {
	switch path {
	case "value":
		return &sf.Value, nil
	case "generatorRef":
		return nil, nil
		// return &sf.GeneratorRef, nil
	default:
		return nil, errors.ErrNoSuchPath
	}
}

func (sf *SecretField) Set(value string) error {
	return errors.ErrCannotSet
}

func (sf *SecretField) Validate() []error {
	var errs []error
	if err := sf.GeneratorRef.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

func (sf *SecretField) GetChildren() []string {
	ret := []string{}
	ret = append(ret, PrefixList("value", sf.Value.GetChildren())...)
	ret = append(ret, PrefixList("generatorRef", sf.GeneratorRef.GetChildren())...)

	return ret
}

func (sf *SecretField) GetValues() map[string]string {
	ret := make(map[string]string)

	for key, val := range PrefixMap("value", sf.Value.GetValues()) {
		ret[key] = val
	}
	for key, val := range PrefixMap("generatorRef", sf.GeneratorRef.GetValues()) {
		ret[key] = val
	}

	return ret
}

func (sf *SecretField) String() string {
	return sf.Value.String()
}
