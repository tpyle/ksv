package types

import "github.com/tpyle/ksv/lib/errors"

type KSVMap map[string]Queryable

func (m *KSVMap) Get(path string) (Queryable, error) {
	if val, ok := (*m)[path]; ok {
		return val, nil
	}
	return nil, errors.ErrNoSuchPath
}

func (m *KSVMap) Set(value string) error {
	return errors.ErrCannotSet
}

func (m *KSVMap) Validate() []error {
	return nil
}

func (m *KSVMap) GetChildren() []string {
	var children []string
	for key := range *m {
		val := (*m)[key]
		children = append(children, PrefixList(key, val.GetChildren())...)
	}
	return children
}

func (m *KSVMap) GetValues() map[string]string {
	values := make(map[string]string)
	for key, val := range *m {
		for subKey, subVal := range PrefixMap(key, val.GetValues()) {
			values[subKey] = subVal
		}
	}
	return values
}

func (m *KSVMap) String() string {
	return ""
}

func (m *KSVMap) ToMap() map[string]string {
	mapped := make(map[string]string)
	for key, val := range *m {
		mapped[key] = val.String()
	}
	return mapped
}

func NewKSVMap(vals map[string]string) KSVMap {
	m := KSVMap{}
	for key, val := range vals {
		val := KSVString(val)
		m[key] = &val
	}
	return m
}
