package types

import (
	"github.com/tpyle/ksv/lib/errors"
)

type Site struct {
	Url       KSVString `json:"url"`
	Name      KSVString `json:"name"`
	AppId     KSVString `json:"appId"`
	GenericId KSVString `json:"genericId"`
	Notes     KSVString `json:"notes"`

	Entries KSVMap `json:"entries"`
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

func (s *Site) Get(path string) (Queryable, error) {
	switch path {
	case "url":
		return &s.Url, nil
	case "name":
		return &s.Name, nil
	case "appId":
		return &s.AppId, nil
	case "genericId":
		return &s.GenericId, nil
	case "notes":
		return &s.Notes, nil
	default:
		if entry, err := s.Entries.Get(path); err == nil {
			return entry, nil
		} else {
			return nil, errors.ErrNoSuchPath
		}
	}
}

func (s *Site) Set(value string) error {
	return errors.ErrCannotSet
}

func (s *Site) GetChildren() []string {
	children := []string{"url", "name", "appId", "genericId", "notes"}
	for key, val := range s.Entries {
		children = append(children, PrefixList(key, val.GetChildren())...)
	}
	return children
}

func (s *Site) GetValues() map[string]string {
	values := map[string]string{
		"url":       s.Url.String(),
		"name":      s.Name.String(),
		"appId":     s.AppId.String(),
		"genericId": s.GenericId.String(),
		"notes":     s.Notes.String(),
	}
	for key, val := range s.Entries {
		entryValues := PrefixMap(key, val.GetValues())
		for k, v := range entryValues {
			values[k] = v
		}
	}
	return values
}

func (s *Site) String() string {
	return ""
}
