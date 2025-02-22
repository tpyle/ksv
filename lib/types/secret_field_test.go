package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/types"
)

func TestGeneratorReference_UnmarshalJSON(t *testing.T) {
	validJSON := `{"ref":"uuid","params":{"length":"10"}}`
	invalidJSON := `{"ref":"invalid","params":{"length":"10"}}`

	t.Run("Valid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(validJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, types.KSVString("uuid"), gr.Ref)
		lval := types.KSVString("10")
		assert.Equal(t, types.KSVMap{"length": &lval}, gr.Params)
		assert.True(t, gr.Supported)
		assert.NotNil(t, gr.Generator)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(invalidJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, types.KSVString("invalid"), gr.Ref)
		lval := types.KSVString("10")
		assert.Equal(t, types.KSVMap{"length": &lval}, gr.Params)
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
				Value: "exampleValue",
				GeneratorRef: types.GeneratorReference{
					Ref: "uuid",
				},
			},
			wantErr: nil,
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
