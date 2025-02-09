package types

import (
	"testing"

	"github.com/tpyle/ksv/lib/errors"
)

func TestIDPValidate(t *testing.T) {
	tests := []struct {
		name    string
		idp     IDP
		wantErr []error
	}{
		{
			name: "valid IDP",
			idp: IDP{
				Url:      "https://example.com",
				Name:     "Example",
				KnownIDP: Google,
			},
			wantErr: nil,
		},
		{
			name: "missing URL",
			idp: IDP{
				Url:      "",
				Name:     "Example",
				KnownIDP: Google,
			},
			wantErr: []error{errors.ErrMissingIDPUrl},
		},
		{
			name: "missing Name",
			idp: IDP{
				Url:      "https://example.com",
				Name:     "",
				KnownIDP: Google,
			},
			wantErr: []error{errors.ErrMissingIDPName},
		},
		{
			name: "invalid KnownIDP",
			idp: IDP{
				Url:      "https://example.com",
				Name:     "Example",
				KnownIDP: "InvalidIDP",
			},
			wantErr: []error{errors.ErrInvalidKnownIDP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.idp.Validate(); !equalErrors(err, tt.wantErr) {
				t.Errorf("IDP.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKnownIDPValidate(t *testing.T) {
	tests := []struct {
		name     string
		knownIDP KnownIDP
		wantErr  []error
	}{
		{
			name:     "valid Google",
			knownIDP: Google,
			wantErr:  nil,
		},
		{
			name:     "valid GitHub",
			knownIDP: GitHub,
			wantErr:  nil,
		},
		{
			name:     "valid Apple",
			knownIDP: Apple,
			wantErr:  nil,
		},
		{
			name:     "valid Generic",
			knownIDP: Generic,
			wantErr:  nil,
		},
		{
			name:     "invalid KnownIDP",
			knownIDP: "InvalidIDP",
			wantErr:  []error{errors.ErrInvalidKnownIDP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.knownIDP.Validate(); !equalErrors(err, tt.wantErr) {
				t.Errorf("KnownIDP.Validate() error = %v, wantErr %v", err, tt.wantErr)
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
