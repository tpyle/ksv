package types

import (
	"encoding/json"

	"github.com/tpyle/ksv/lib/errors"
	"github.com/tpyle/ksv/lib/generators"
)

type GeneratorReference struct {
	Ref       string               `json:"ref"`
	Params    map[string]string    `json:"params"`
	Generator generators.Generator `json:"-"`
	Supported bool                 `json:"-"`
}

func (gr *GeneratorReference) UnmarshalJSON(data []byte) error {
	type Alias GeneratorReference
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(gr),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	gr.Params = aux.Params
	gr.Ref = aux.Ref
	gr.Supported, gr.Generator = generators.GetGenerator(gr.Ref)
	return nil
}

func (gr *GeneratorReference) Validate() []error {
	return nil
}

type SecretField struct {
	FieldName    string             `json:"field"`
	Value        string             `json:"value"`
	GeneratorRef GeneratorReference `json:"generatorRef"`
}

func (sf *SecretField) Validate() []error {
	var errs []error
	if sf.FieldName == "" {
		errs = append(errs, errors.ErrMissingSecretFieldName)
	}
	if err := sf.GeneratorRef.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}
