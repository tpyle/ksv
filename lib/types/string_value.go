package types

import "github.com/tpyle/ksv/lib/errors"

type KSVString string

func (s *KSVString) Get(path string) (Queryable, error) {
	if path != "" {
		return nil, errors.ErrInvalidPath
	}
	return s, nil
}

func (s *KSVString) Set(value string) error {
	*s = KSVString(value)
	return nil
}

func (s *KSVString) Validate() []error {
	return nil
}

func (s *KSVString) GetChildren() []string {
	return []string{""}
}

func (s *KSVString) GetValues() map[string]string {
	return map[string]string{"": s.String()}
}

func (s *KSVString) String() string {
	return string(*s)
}
