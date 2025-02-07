package types

type CustomField struct {
	FieldName string `json:"field"`
	Value     string `json:"value"`
}

func (cf *CustomField) Validate() error {
	return nil
}
