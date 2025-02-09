package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/errors"
	"github.com/tpyle/ksv/lib/types"
)

func TestGeneratorReference_UnmarshalJSON(t *testing.T) {
	validJSON := `{"ref":"uuid","params":{"length":"10"}}`
	invalidJSON := `{"ref":"invalid","params":{"length":"10"}}`

	t.Run("Valid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(validJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, "uuid", gr.Ref)
		assert.Equal(t, map[string]string{"length": "10"}, gr.Params)
		assert.True(t, gr.Supported)
		assert.NotNil(t, gr.Generator)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(invalidJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, "invalid", gr.Ref)
		assert.Equal(t, map[string]string{"length": "10"}, gr.Params)
		assert.False(t, gr.Supported)
		assert.NotNil(t, gr.Generator)
	})
}

func TestSecretField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   types.SecretField
		wantErr []error
	}{
		{
			name: "valid SecretField",
			field: types.SecretField{
				FieldName: "exampleField",
				Value:     "exampleValue",
				GeneratorRef: types.GeneratorReference{
					Ref:    "uuid",
					Params: map[string]string{"length": "10"},
				},
			},
			wantErr: nil,
		},
		{
			name: "missing FieldName",
			field: types.SecretField{
				FieldName: "",
				Value:     "exampleValue",
				GeneratorRef: types.GeneratorReference{
					Ref:    "uuid",
					Params: map[string]string{"length": "10"},
				},
			},
			wantErr: []error{errors.ErrMissingSecretFieldName},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); !equalErrors(err, tt.wantErr) {
				t.Errorf("SecretField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func equalErrors(got, want []error) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
