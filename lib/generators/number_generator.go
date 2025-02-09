package generators

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/go-viper/mapstructure/v2"
	"github.com/tpyle/ksv/lib/errors"
)

const (
	NumberGeneratorRef = "number"
	NumberCharacters   = "0123456789"
)

type NumberGenerator struct {
}

type NumberGeneratorParams struct {
	Length int `json:"length"`
}

func DecodeNumberGeneratorParams(params map[string]string) (NumberGeneratorParams, error) {
	var numberGeneratorParams struct {
		Length string `json:"length"`
	}
	err := mapstructure.Decode(params, &numberGeneratorParams)

	length := 0
	if numberGeneratorParams.Length != "" {
		length, err = strconv.Atoi(numberGeneratorParams.Length)
		if err != nil {
			return NumberGeneratorParams{}, err
		}
	} else {
		length = 0
	}
	if err != nil {
		return NumberGeneratorParams{}, err
	}
	return NumberGeneratorParams{
		Length: length,
	}, nil
}

func (n *NumberGenerator) Generate(params map[string]string) (string, error) {
	numberGeneratorParams, err := DecodeNumberGeneratorParams(params)
	if err != nil {
		return "", err
	}

	if numberGeneratorParams.Length == 0 {
		return "", errors.ErrGeneratorInvalidLength
	}

	randReader := rand.Reader

	length := numberGeneratorParams.Length
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(randReader, big.NewInt(int64(len(NumberCharacters))))
		if err != nil {
			return "", err
		}
		result[i] = NumberCharacters[n.Int64()]
	}

	return string(result), nil
}

func (n *NumberGenerator) GetRef() string {
	return NumberGeneratorRef
}
