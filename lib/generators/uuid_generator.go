package generators

import "github.com/google/uuid"

const (
	UUIDGeneratorRef = "uuid"
)

type UUIDGenerator struct {
}

type UUIDGeneratorParams struct {
}

func DecodeUUIDGeneratorParams(params map[string]string) (UUIDGeneratorParams, error) {
	return UUIDGeneratorParams{}, nil
}

func (n *UUIDGenerator) Generate(params map[string]string) (string, error) {
	_, err := DecodeUUIDGeneratorParams(params)
	if err != nil {
		return "", err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func (n *UUIDGenerator) GetRef() string {
	return UUIDGeneratorRef
}

func (n *UUIDGenerator) GetDefaultParams() map[string]string {
	return map[string]string{}
}
