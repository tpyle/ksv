package types

import "github.com/tpyle/ksv/lib/errors"

type Entry struct {
	DistinguishingName string `json:"distinguishingName"` // This is the unique identifier for the entry if provided
	IsIDP              bool   `json:"isIDP"`              // This is true if the entry is an IDP
	IDP                IDP    `json:"idp"`                // This is the IDP object if the entry is an IDP
	Username           string `json:"username"`           // This is the username for the entry
	Email              string `json:"email"`              // This is the email for the entry
	Notes              string `json:"notes"`              // This is the notes for the entry

	CustomFields []CustomField `json:"customFields"`
	SecretFields []SecretField `json:"secretFields"`
}

func (e *Entry) Validate() []error {
	var errs []error

	if e.DistinguishingName == "" {
		errs = append(errs, errors.ErrMissingDistinguishingName)
	}

	if e.IsIDP {
		err := e.IDP.Validate()
		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	for _, cf := range e.CustomFields {
		err := cf.Validate()
		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	for _, sf := range e.SecretFields {
		err := sf.Validate()
		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	return errs
}
