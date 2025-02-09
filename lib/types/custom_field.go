package types

import "github.com/tpyle/ksv/lib/errors"

type CustomField struct {
	FieldName string `json:"field"`
	Value     string `json:"value"`
}

func (cf *CustomField) Validate() []error {
	var errs []error
	if cf.FieldName == "" {
		errs = append(errs, errors.ErrMissingCustomFieldName)
	}
	return errs
}
