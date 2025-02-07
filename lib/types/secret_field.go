package types

type SecretField struct {
	FieldName       string            `json:"field"`
	Value           string            `json:"value"`
	GeneratorRef    string            `json:"generatorRef"`
	GeneratorParams map[string]string `json:"generatorParams"`
}

func (sf *SecretField) Validate() error {
	return nil
}
