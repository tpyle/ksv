package types

import "github.com/tpyle/ksv/lib/ksverrors"

type Entry struct {
	IsIDP    *KSVBool  `json:"isIDP"`    // This is true if the entry is an IDP
	IDP      IDP       `json:"idp"`      // This is the IDP object if the entry is an IDP
	Username KSVString `json:"username"` // This is the username for the entry
	Email    KSVString `json:"email"`    // This is the email for the entry
	Notes    KSVString `json:"notes"`    // This is the notes for the entry

	CustomFields KSVMap[*KSVString] `json:"customFields"`
	SecretFields KSVMap[*KSVString] `json:"secretFields"`
}

func (e *Entry) Get(path string) (Queryable, error) {
	switch path {
	case "isIDP":
		return e.IsIDP, nil
	case "idp":
		return &e.IDP, nil
	case "username":
		return &e.Username, nil
	case "email":
		return &e.Email, nil
	case "notes":
		return &e.Notes, nil
	default:
		if cf, err := e.CustomFields.Get(path); err == nil {
			return cf, nil
		} else if sf, err := e.SecretFields.Get(path); err == nil {
			return sf, nil
		} else {
			return nil, ksverrors.ErrNoSuchPath
		}
	}
}

func (e *Entry) Set(value string) error {
	return ksverrors.ErrCannotSet
}

func (e *Entry) Validate() []error {
	var errs []error

	if e.Username.String() == "" {
		errs = append(errs, ksverrors.ErrMissingUsername)
	}

	if e.IsIDP.ToBool() {
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

func (e *Entry) GetChildren() []string {
	children := []string{"isIDP", "username", "email", "notes"}

	children = append(children, PrefixList("idp", PrefixList("idp", e.IDP.GetChildren()))...)

	for key, val := range e.CustomFields {
		children = append(children, PrefixList(key, val.GetChildren())...)
	}

	for key, val := range e.SecretFields {
		children = append(children, PrefixList(key, val.GetChildren())...)
	}

	return children
}

func (e *Entry) GetValues() map[string]string {
	values := map[string]string{
		"isIDP":    e.IsIDP.String(),
		"username": e.Username.String(),
		"email":    e.Email.String(),
		"notes":    e.Notes.String(),
	}

	idpValues := PrefixMap("idp", e.IDP.GetValues())
	for k, v := range idpValues {
		values[k] = v
	}

	for key, val := range e.CustomFields {
		cfValues := PrefixMap(key, val.GetValues())
		for k, v := range cfValues {
			values[k] = v
		}
	}

	for key, val := range e.SecretFields {
		sfValues := PrefixMap(key, val.GetValues())
		for k, v := range sfValues {
			values[k] = v
		}
	}

	return values
}

func (e *Entry) String() string {
	return ""
}
