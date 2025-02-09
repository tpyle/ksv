package generators

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/go-viper/mapstructure/v2"
	"github.com/tpyle/ksv/lib/errors"
)

const (
	OWASPGeneratorRef      = "password"
	LowercaseCharacters    = "abcdefghijklmnopqrstuvwxyz"
	UppercaseCharacters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	OWASPSpecialCharacters = " !\"#$%&'()*+,-./:;<=>?@[\\]&_`{|}~"

	owaspLowercaseRef = "l"
	owaspUppercaseRef = "u"
	owaspNumberRef    = "n"
	owaspSpecialRef   = "s"
	owaspAnyRef       = "x"
)

type OWASPGenerator struct {
}

type OWASPGeneratorParams struct {
	Length       int    `mapstructure:"length" json:"length"`
	MinLowercase int    `mapstructure:"min_lowercase" json:"min_lowercase"`
	MinUppercase int    `mapstructure:"min_uppercase" json:"min_uppercase"`
	MinNumbers   int    `mapstructure:"min_numbers" json:"min_numbers"`
	MinSpecial   int    `mapstructure:"min_special" json:"min_special"`
	SpecialChars string `mapstructure:"special_chars" json:"special_chars"`
}

func DecodeOWASPGeneratorParams(mapParams map[string]string) (OWASPGeneratorParams, error) {
	var owaspGeneratorParams struct {
		Length       string `mapstructure:"length" json:"length"`
		MinLowercase string `mapstructure:"min_lowercase" json:"min_lowercase"`
		MinUppercase string `mapstructure:"min_uppercase" json:"min_uppercase"`
		MinNumbers   string `mapstructure:"min_numbers" json:"min_numbers"`
		MinSpecial   string `mapstructure:"min_special" json:"min_special"`
		SpecialChars string `mapstructure:"special_chars" json:"special_chars"`
	}
	err := mapstructure.Decode(mapParams, &owaspGeneratorParams)
	if err != nil {
		return OWASPGeneratorParams{}, err
	}

	var params OWASPGeneratorParams
	if owaspGeneratorParams.Length != "" {
		params.Length, err = strconv.Atoi(owaspGeneratorParams.Length)
		if err != nil {
			return params, err
		}
	}
	if owaspGeneratorParams.MinLowercase != "" {
		params.MinLowercase, err = strconv.Atoi(owaspGeneratorParams.MinLowercase)
		if err != nil {
			return params, err
		}
	}
	if owaspGeneratorParams.MinUppercase != "" {
		params.MinUppercase, err = strconv.Atoi(owaspGeneratorParams.MinUppercase)
		if err != nil {
			return params, err
		}
	}
	if owaspGeneratorParams.MinNumbers != "" {
		params.MinNumbers, err = strconv.Atoi(owaspGeneratorParams.MinNumbers)
		if err != nil {
			return params, err
		}
	}
	if owaspGeneratorParams.MinSpecial != "" {
		params.MinSpecial, err = strconv.Atoi(owaspGeneratorParams.MinSpecial)
		if err != nil {
			return params, err
		}
	}
	if owaspGeneratorParams.SpecialChars != "" {
		params.SpecialChars = owaspGeneratorParams.SpecialChars
	} else {
		params.SpecialChars = OWASPSpecialCharacters
	}
	return params, nil
}

func ShuffleArray(array []string) []string {
	arrayLen := big.NewInt(int64(len(array)))
	for i := range array {
		j, err := rand.Int(rand.Reader, arrayLen)
		if err != nil {
			panic(err)
		}
		array[i], array[j.Int64()] = array[j.Int64()], array[i]
	}
	return array
}

func CreateOrder(params OWASPGeneratorParams) []string {
	var order []string
	for i := 0; i < params.MinLowercase; i++ {
		order = append(order, owaspLowercaseRef)
	}
	for i := 0; i < params.MinUppercase; i++ {
		order = append(order, owaspUppercaseRef)
	}
	for i := 0; i < params.MinNumbers; i++ {
		order = append(order, owaspNumberRef)
	}
	for i := 0; i < params.MinSpecial; i++ {
		order = append(order, owaspSpecialRef)
	}
	iterations := params.Length - len(order)
	for i := 0; i < iterations; i++ {
		order = append(order, owaspAnyRef)
	}
	return ShuffleArray(order)
}

func (n *OWASPGenerator) Generate(params map[string]string) (string, error) {
	owaspGeneratorParams, err := DecodeOWASPGeneratorParams(params)
	if err != nil {
		return "", err
	}

	if owaspGeneratorParams.Length == 0 {
		return "", errors.ErrGeneratorInvalidLength
	}

	if owaspGeneratorParams.MinLowercase+owaspGeneratorParams.MinUppercase+owaspGeneratorParams.MinNumbers+owaspGeneratorParams.MinSpecial > owaspGeneratorParams.Length {
		return "", errors.ErrGeneratorImpossibleParams
	}

	allCharacters := LowercaseCharacters + UppercaseCharacters + owaspGeneratorParams.SpecialChars + NumberCharacters
	randReader := rand.Reader
	length := owaspGeneratorParams.Length
	result := make([]byte, length)

	order := CreateOrder(owaspGeneratorParams)

	for i := 0; i < length; i++ {
		switch order[i] {
		case owaspLowercaseRef:
			n, err := rand.Int(randReader, big.NewInt(int64(len(LowercaseCharacters))))
			if err != nil {
				return "", err
			}
			result[i] = LowercaseCharacters[n.Int64()]
		case owaspUppercaseRef:
			n, err := rand.Int(randReader, big.NewInt(int64(len(UppercaseCharacters))))
			if err != nil {
				return "", err
			}
			result[i] = UppercaseCharacters[n.Int64()]
		case owaspNumberRef:
			n, err := rand.Int(randReader, big.NewInt(int64(len(NumberCharacters))))
			if err != nil {
				return "", err
			}
			result[i] = NumberCharacters[n.Int64()]
		case owaspSpecialRef:
			n, err := rand.Int(randReader, big.NewInt(int64(len(owaspGeneratorParams.SpecialChars))))
			if err != nil {
				return "", err
			}
			result[i] = owaspGeneratorParams.SpecialChars[n.Int64()]
		case owaspAnyRef:
			n, err := rand.Int(randReader, big.NewInt(int64(len(allCharacters))))
			if err != nil {
				return "", err
			}
			result[i] = allCharacters[n.Int64()]
		}
	}

	return string(result), nil
}

func (n *OWASPGenerator) GetRef() string {
	return OWASPGeneratorRef
}
