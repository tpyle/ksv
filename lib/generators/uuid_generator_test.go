package generators

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDecodeUUIDGeneratorParams(t *testing.T) {
	t.Run("Valid params", func(t *testing.T) {
		params := map[string]string{}
		result, err := DecodeUUIDGeneratorParams(params)
		assert.NoError(t, err)
		assert.Equal(t, UUIDGeneratorParams{}, result)
	})
}

func TestUUIDGenerator_Generate(t *testing.T) {
	generator := &UUIDGenerator{}

	t.Run("Generate UUID", func(t *testing.T) {
		params := map[string]string{}
		result, err := generator.Generate(params)
		assert.NoError(t, err)

		_, err = uuid.Parse(result)
		assert.NoError(t, err, "Generated string is not a valid UUID")
	})
}
