package generators

import (
	"testing"

	"github.com/tpyle/ksv/lib/ksverrors"
)

func TestNumberGenerate(t *testing.T) {
	generator := &NumberGenerator{}

	t.Run("Length is zero", func(t *testing.T) {
		params := map[string]string{"length": "0"}
		_, err := generator.Generate(params)
		if err != ksverrors.ErrGeneratorInvalidLength {
			t.Errorf("expected error %v, got %v", ksverrors.ErrGeneratorInvalidLength, err)
		}
	})

	t.Run("Valid length", func(t *testing.T) {
		params := map[string]string{"length": "5"}
		result, err := generator.Generate(params)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(result) != 5 {
			t.Errorf("expected length 5, got %d", len(result))
		}
		for _, char := range result {
			if !contains(NumberCharacters, byte(char)) {
				t.Errorf("unexpected character: %v", char)
			}
		}
	})

	t.Run("Invalid params", func(t *testing.T) {
		params := map[string]string{"invalid": "param"}
		_, err := generator.Generate(params)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
