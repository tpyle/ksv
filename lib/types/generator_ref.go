package types

import (
	"encoding/json"

	"github.com/tpyle/ksv/lib/errors"
	"github.com/tpyle/ksv/lib/generators"
)

type GeneratorReference struct {
	Ref       KSVString            `json:"ref"`
	Params    KSVMap               `json:"params"`
	Generator generators.Generator `json:"-"`
	Supported bool                 `json:"-"`
}

func (gr *GeneratorReference) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Ref    KSVString
		Params map[string]string
	}
	var aux = Alias{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	gr.Params = NewKSVMap(aux.Params)
	gr.Ref = aux.Ref
	gr.Supported, gr.Generator = generators.GetGenerator(gr.Ref.String())
	return nil
}

func (gr *GeneratorReference) Get(path string) (Queryable, error) {
	switch path {
	case "ref":
		return &gr.Ref, nil
	case "params":
		return &gr.Params, nil
	default:
		return nil, errors.ErrNoSuchPath
	}
}

func (gr *GeneratorReference) Set(value string) error {
	return errors.ErrCannotSet
}

func (gr *GeneratorReference) Validate() []error {
	return nil
}

func (gr *GeneratorReference) GetChildren() []string {
	ret := []string{}
	ret = append(ret, PrefixList("ref", gr.Ref.GetChildren())...)
	ret = append(ret, PrefixList("params", gr.Params.GetChildren())...)

	return ret
}

func (gr *GeneratorReference) GetValues() map[string]string {
	ret := make(map[string]string)

	for key, val := range PrefixMap("ref", gr.Ref.GetValues()) {
		ret[key] = val
	}
	for key, val := range PrefixMap("params", gr.Params.GetValues()) {
		ret[key] = val
	}

	return ret
}
