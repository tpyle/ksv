package types

import "github.com/tpyle/ksv/lib/errors"

type IDP struct {
	Url      string   `json:"url"`
	Name     string   `json:"name"`
	KnownIDP KnownIDP `json:"knownIDP"`
}

func (i *IDP) Validate() error {
	if i.Url == "" {
		return errors.ErrMissingIDPUrl
	}

	if i.Name == "" {
		return errors.ErrMissingIDPName
	}

	if err := i.KnownIDP.Validate(); err != nil {
		return err
	}

	return nil
}

type KnownIDP string

const (
	Google  KnownIDP = "Google"
	GitHub  KnownIDP = "GitHub"
	Apple   KnownIDP = "Apple"
	Generic KnownIDP = "Generic"
)

func (k KnownIDP) Validate() error {
	switch k {
	case Google:
		return nil
	case GitHub:
		return nil
	case Apple:
		return nil
	case Generic:
		return nil
	default:
		return errors.ErrInvalidKnownIDP
	}
}
