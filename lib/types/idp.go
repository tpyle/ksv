package types

import "github.com/tpyle/ksv/lib/ksverrors"

type IDP struct {
	Url      KSVString `json:"url"`
	Name     KSVString `json:"name"`
	KnownIDP KnownIDP  `json:"knownIDP"`
}

func (i *IDP) Get(pathElement string) (Queryable, error) {
	switch pathElement {
	case "url":
		return &i.Url, nil
	case "name":
		return &i.Name, nil
	case "knownIDP":
		return &i.KnownIDP, nil
	default:
		return nil, ksverrors.ErrNoSuchPath
	}
}

func (i *IDP) Validate() []error {
	var errs []error
	if i.Url.Value == "" {
		errs = append(errs, ksverrors.ErrMissingIDPUrl)
	}

	if i.Name.Value == "" {
		errs = append(errs, ksverrors.ErrMissingIDPName)
	}

	if err := i.KnownIDP.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

func (i *IDP) Set(value string) error {
	return ksverrors.ErrCannotSet
}

func (i *IDP) GetChildren() []string {
	return []string{"url", "name", "knownIDP"}
}

func (i *IDP) GetValues() map[string]string {
	return map[string]string{
		"url":      i.Url.String(),
		"name":     i.Name.String(),
		"knownIDP": i.KnownIDP.String(),
	}
}

func (i *IDP) String() string {
	return ""
}

type KnownIDP KSVString

var (
	Google = KnownIDP{
		Value: "Google",
	}
	GitHub = KnownIDP{
		Value: "GitHub",
	}
	Apple = KnownIDP{
		Value: "Apple",
	}
	Generic = KnownIDP{
		Value: "Generic",
	}
)

func (k *KnownIDP) Get(path string) (Queryable, error) {
	if len(path) > 0 {
		return nil, ksverrors.ErrNoSuchPath
	}
	return k, nil
}

func (k *KnownIDP) Set(value string) error {
	switch value {
	case "Google":
		*k = Google
	case "GitHub":
		*k = GitHub
	case "Apple":
		*k = Apple
	case "Generic":
		*k = Generic
	default:
		return ksverrors.ErrInvalidKnownIDP
	}
	return nil
}

func (k *KnownIDP) Validate() []error {
	switch k.Value {
	case Google.Value:
		return nil
	case GitHub.Value:
		return nil
	case Apple.Value:
		return nil
	case Generic.Value:
		return nil
	default:
		return []error{ksverrors.ErrInvalidKnownIDP}
	}
}

func (k *KnownIDP) GetChildren() []string {
	return nil
}

func (k *KnownIDP) GetValues() map[string]string {
	return nil
}

func (k *KnownIDP) String() string {
	return k.Value
}
