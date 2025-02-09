package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/errors"
)

func TestDecodeOWASPGeneratorParams(t *testing.T) {
	t.Run("Valid length", func(t *testing.T) {
		params := map[string]string{"length": "10"}
		result, err := DecodeOWASPGeneratorParams(params)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.Length != 10 {
			t.Errorf("expected length 10, got %d", result.Length)
		}
	})

	t.Run("Invalid length", func(t *testing.T) {
		params := map[string]string{"length": "invalid"}
		_, err := DecodeOWASPGeneratorParams(params)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Missing length", func(t *testing.T) {
		params := map[string]string{}
		result, err := DecodeOWASPGeneratorParams(params)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.Length != 0 {
			t.Errorf("expected length 0, got %d", result.Length)
		}
	})

	t.Run("Valid all params", func(t *testing.T) {
		params := map[string]string{
			"length":        "12",
			"min_lowercase": "2",
			"min_uppercase": "2",
			"min_numbers":   "2",
			"min_special":   "2",
			"special_chars": "@#$%",
		}
		result, err := DecodeOWASPGeneratorParams(params)
		assert.NoError(t, err)
		assert.Equal(t, 12, result.Length)
		assert.Equal(t, 2, result.MinLowercase)
		assert.Equal(t, 2, result.MinUppercase)
		assert.Equal(t, 2, result.MinNumbers)
		assert.Equal(t, 2, result.MinSpecial)
		assert.Equal(t, "@#$%", result.SpecialChars)
	})

	t.Run("Invalid min_lowercase", func(t *testing.T) {
		params := map[string]string{"min_lowercase": "invalid"}
		_, err := DecodeOWASPGeneratorParams(params)
		assert.Error(t, err)
	})

	t.Run("Invalid min_uppercase", func(t *testing.T) {
		params := map[string]string{"min_uppercase": "invalid"}
		_, err := DecodeOWASPGeneratorParams(params)
		assert.Error(t, err)
	})

	t.Run("Invalid min_numbers", func(t *testing.T) {
		params := map[string]string{"min_numbers": "invalid"}
		_, err := DecodeOWASPGeneratorParams(params)
		assert.Error(t, err)
	})

	t.Run("Invalid min_special", func(t *testing.T) {
		params := map[string]string{"min_special": "invalid"}
		_, err := DecodeOWASPGeneratorParams(params)
		assert.Error(t, err)
	})

	t.Run("Invalid special_chars", func(t *testing.T) {
		params := map[string]string{"special_chars": ""}
		result, err := DecodeOWASPGeneratorParams(params)
		assert.NoError(t, err)
		assert.Equal(t, OWASPSpecialCharacters, result.SpecialChars)
	})
}

func TestOWASPGenerator_Generate(t *testing.T) {
	generator := &OWASPGenerator{}

	t.Run("valid length", func(t *testing.T) {
		params := map[string]string{"length": "10"}
		result, err := generator.Generate(params)
		assert.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("invalid length", func(t *testing.T) {
		params := map[string]string{"length": "invalid"}
		_, err := generator.Generate(params)
		assert.Error(t, err)
	})

	t.Run("zero length", func(t *testing.T) {
		params := map[string]string{"length": "0"}
		_, err := generator.Generate(params)
		assert.ErrorIs(t, err, errors.ErrGeneratorInvalidLength)
	})

	t.Run("missing length", func(t *testing.T) {
		params := map[string]string{}
		_, err := generator.Generate(params)
		assert.ErrorIs(t, err, errors.ErrGeneratorInvalidLength)
	})

	t.Run("valid all params", func(t *testing.T) {
		params := map[string]string{
			"length":        "12",
			"min_lowercase": "2",
			"min_uppercase": "2",
			"min_numbers":   "2",
			"min_special":   "2",
			"special_chars": "@#$%",
		}
		result, err := generator.Generate(params)
		assert.NoError(t, err)
		assert.Len(t, result, 12)
	})

	t.Run("impossible params", func(t *testing.T) {
		params := map[string]string{
			"length":        "5",
			"min_lowercase": "2",
			"min_uppercase": "2",
			"min_numbers":   "2",
			"min_special":   "2",
		}
		_, err := generator.Generate(params)
		assert.ErrorIs(t, err, errors.ErrGeneratorImpossibleParams)
	})
}
