package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/ksverrors"
)

func TestUnsuportedGenerator_Generate(t *testing.T) {
	generator := &UnsupportedGenerator{}

	t.Run("Try Generate Unsupported", func(t *testing.T) {
		params := map[string]string{}
		_, err := generator.Generate(params)
		assert.ErrorIs(t, err, ksverrors.ErrUnsupportedGenerator)
	})
}
