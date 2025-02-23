package types

import "github.com/tpyle/ksv/lib/ksverrors"

type KSVMap[T Queryable] map[string]T

func (m *KSVMap[T]) Get(path string) (Queryable, error) {
	if val, ok := (*m)[path]; ok {
		return val, nil
	}
	var zero T
	return zero, ksverrors.ErrNoSuchPath
}

func (m *KSVMap[T]) Set(value string) error {
	return ksverrors.ErrCannotSet
}

func (m *KSVMap[T]) Validate() []error {
	return nil
}

func (m *KSVMap[T]) GetChildren() []string {
	var children []string
	for key := range *m {
		val := (*m)[key]
		children = append(children, PrefixList(key, val.GetChildren())...)
	}
	return children
}

func (m *KSVMap[T]) GetValues() map[string]string {
	values := make(map[string]string)
	for key, val := range *m {
		for subKey, subVal := range PrefixMap(key, val.GetValues()) {
			values[subKey] = subVal
		}
	}
	return values
}

func (m *KSVMap[T]) String() string {
	return ""
}

func (m *KSVMap[T]) ToMap() map[string]string {
	mapped := make(map[string]string)
	for key, val := range *m {
		mapped[key] = val.String()
	}
	return mapped
}

func NewKSVMap[T Queryable](vals map[string]string, createFunc func(string) T) *KSVMap[T] {
	m := KSVMap[T]{}
	for key, val := range vals {
		m[key] = createFunc(val)
	}
	return &m
}
