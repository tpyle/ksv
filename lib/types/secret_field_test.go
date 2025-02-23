package types_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/types"
)

func TestGeneratorReference_UnmarshalJSON(t *testing.T) {
	validJSON := `{"ref":{"value":"uuid"},"params":{"length":"10"}}`
	invalidJSON := `{"ref":{"value":"invalid"},"params":{"length":"10"}}`

	t.Run("Valid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(validJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, &types.KSVString{
			Value: "uuid"}, gr.Ref)
		lval := types.NewKSVString("10")
		assert.Equal(t, &types.KSVMap[*types.KSVString]{"length": lval}, gr.Params)
		assert.True(t, gr.Supported)
		assert.NotNil(t, gr.Generator)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		var gr types.GeneratorReference
		err := json.Unmarshal([]byte(invalidJSON), &gr)
		assert.NoError(t, err)
		assert.Equal(t, types.NewKSVString("invalid"), gr.Ref)
		lval := types.NewKSVString("10")
		assert.Equal(t, &types.KSVMap[*types.KSVString]{"length": lval}, gr.Params)
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
				Value: types.NewKSVString("exampleValue"),
				GeneratorRef: &types.GeneratorReference{
					Ref: types.NewKSVString("uuid"),
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
