package types

import "github.com/tpyle/ksv/lib/errors"

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
		return nil, errors.ErrNoSuchPath
	}
}

func (i *IDP) Validate() []error {
	var errs []error
	if i.Url == "" {
		errs = append(errs, errors.ErrMissingIDPUrl)
	}

	if i.Name == "" {
		errs = append(errs, errors.ErrMissingIDPName)
	}

	if err := i.KnownIDP.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

func (i *IDP) Set(value string) error {
	return errors.ErrCannotSet
}

func (i *IDP) GetChildren() []string {
	return []string{"url", "name", "knownIDP"}
}

func (i *IDP) GetValues() map[string]string {
	return map[string]string{
		"url":      string(i.Url),
		"name":     string(i.Name),
		"knownIDP": string(i.KnownIDP),
	}
}

func (i *IDP) String() string {
	return string(i.Name)
}

type KnownIDP string

const (
	Google  KnownIDP = "Google"
	GitHub  KnownIDP = "GitHub"
	Apple   KnownIDP = "Apple"
	Generic KnownIDP = "Generic"
)

func (k *KnownIDP) Get(path string) (Queryable, error) {
	if len(path) > 0 {
		return nil, errors.ErrNoSuchPath
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
		return errors.ErrInvalidKnownIDP
	}
	return nil
}

func (k *KnownIDP) Validate() []error {
	switch *k {
	case Google:
		return nil
	case GitHub:
		return nil
	case Apple:
		return nil
	case Generic:
		return nil
	default:
		return []error{errors.ErrInvalidKnownIDP}
	}
}

func (k *KnownIDP) GetChildren() []string {
	return nil
}

func (k *KnownIDP) GetValues() map[string]string {
	return nil
}

func (k *KnownIDP) String() string {
	return string(*k)
}
