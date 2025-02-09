package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/errors"
)

func TestSite_Validate(t *testing.T) {
	tests := []struct {
		name    string
		site    Site
		wantErr bool
		errors  []error
	}{
		{
			name: "Valid site",
			site: Site{
				Url:   "https://example.com",
				Name:  "Example",
				AppId: "app123",
			},
			wantErr: false,
			errors:  nil,
		},
		{
			name: "Missing Url, AppId, and GenericId",
			site: Site{
				Name: "Example",
			},
			wantErr: true,
			errors:  []error{errors.ErrMissingSiteUrlOrAppIdOrGenericId},
		},
		{
			name: "Missing Name",
			site: Site{
				Url:   "https://example.com",
				AppId: "app123",
			},
			wantErr: true,
			errors:  []error{errors.ErrMissingSiteName},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.site.Validate()
			if tt.wantErr {
				assert.NotEmpty(t, errs)
				for _, err := range tt.errors {
					assert.Contains(t, errs, err)
				}
			} else {
				assert.Empty(t, errs)
			}
		})
	}
}
