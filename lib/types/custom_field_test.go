package types

import (
	"testing"

	"github.com/tpyle/ksv/lib/errors"
)

func TestCustomFieldValidate(t *testing.T) {
	tests := []struct {
		name    string
		field   CustomField
		wantErr []error
	}{
		{
			name: "valid CustomField",
			field: CustomField{
				FieldName: "exampleField",
				Value:     "exampleValue",
			},
			wantErr: nil,
		},
		{
			name: "missing FieldName",
			field: CustomField{
				FieldName: "",
				Value:     "exampleValue",
			},
			wantErr: []error{errors.ErrMissingCustomFieldName},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); !equalErrors(err, tt.wantErr) {
				t.Errorf("CustomField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
