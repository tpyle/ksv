package types

type Entry struct {
	DistinguishingName string `json:"distinguishingName"` // This is the unique identifier for the entry if provided
	IsIDP              bool   `json:"isIDP"`              // This is true if the entry is an IDP
	IDP                IDP    `json:"idp"`                // This is the IDP object if the entry is an IDP
	Username           string `json:"username"`           // This is the username for the entry
	Email              string `json:"email"`              // This is the email for the entry

	CustomFields []CustomField `json:"customFields"`
	SecretFields []SecretField `json:"secretFields"`
}

func (e *Entry) Validate() error {
	if e.IsIDP {
		if err := e.IDP.Validate(); err != nil {
			return err
		}
	}

	for _, cf := range e.CustomFields {
		if err := cf.Validate(); err != nil {
			return err
		}
	}

	for _, sf := range e.SecretFields {
		if err := sf.Validate(); err != nil {
			return err
		}
	}

	return nil
}
