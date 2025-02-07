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
}
