package types

import "github.com/tpyle/ksv/lib/ksverrors"

type KSVString struct {
	Value string
}

var (
	KSVStringEmpty = KSVString{Value: ""}
)

func (s *KSVString) Get(path string) (Queryable, error) {
	if path != "" {
		return nil, ksverrors.ErrInvalidPath
	}
	return s, nil
}

func (s *KSVString) Set(value string) error {
	s.Value = value
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
	return string(s.Value)
}

func (s *KSVString) IsEmpty() bool {
	return s == nil || s.Value == ""
}

func NewKSVString(value string) *KSVString {
	return &KSVString{Value: value}
}
