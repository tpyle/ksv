package types

import "github.com/tpyle/ksv/lib/errors"

type Site struct {
	Url       string `json:"url"`
	Name      string `json:"name"`
	AppId     string `json:"appId"`
	GenericId string `json:"genericId"`
	Notes     string `json:"notes"`

	Entries []Entry `json:"entries"`
}

func (s *Site) Validate() []error {
	var errs []error
	if s.Url == "" && s.AppId == "" && s.GenericId == "" {
		errs = append(errs, errors.ErrMissingSiteUrlOrAppIdOrGenericId)
	}

	if s.Name == "" {
		errs = append(errs, errors.ErrMissingSiteName)
	}

	for _, e := range s.Entries {
		errs = append(errs, e.Validate()...)
	}

	return errs
}
