package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/ksverrors"
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
				Url:   NewKSVString("https://example.com"),
				Name:  NewKSVString("Example"),
				AppId: NewKSVString("app123"),
			},
			wantErr: false,
			errors:  nil,
		},
		{
			name: "Missing Url, AppId, and GenericId",
			site: Site{
				Name: NewKSVString("Example"),
			},
			wantErr: true,
			errors:  []error{ksverrors.ErrMissingSiteUrlOrAppIdOrGenericId},
		},
		{
			name: "Missing Name",
			site: Site{
				Url:   NewKSVString("https://example.com"),
				AppId: NewKSVString("app123"),
			},
			wantErr: true,
			errors:  []error{ksverrors.ErrMissingSiteName},
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

func TestSite_Get(t *testing.T) {
	site := Site{
		Url:       NewKSVString("https://example.com"),
		Name:      NewKSVString("Example"),
		AppId:     NewKSVString("app123"),
		GenericId: NewKSVString("gen123"),
		Notes:     NewKSVString("Some notes"),
		Entries:   KSVMap[*Entry]{},
	}

	tests := []struct {
		name    string
		path    string
		want    Queryable
		wantErr bool
	}{
		{
			name:    "Get Url",
			path:    "url",
			want:    site.Url,
			wantErr: false,
		},
		{
			name:    "Get Name",
			path:    "name",
			want:    site.Name,
			wantErr: false,
		},
		{
			name:    "Get AppId",
			path:    "appId",
			want:    site.AppId,
			wantErr: false,
		},
		{
			name:    "Get GenericId",
			path:    "genericId",
			want:    site.GenericId,
			wantErr: false,
		},
		{
			name:    "Get Notes",
			path:    "notes",
			want:    site.Notes,
			wantErr: false,
		},
		{
			name:    "Invalid Path",
			path:    "invalid",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := site.Get(tt.path)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSite_Set(t *testing.T) {
	site := Site{}
	err := site.Set("value")
	assert.Error(t, err)
	assert.Equal(t, ksverrors.ErrCannotSet, err)
}

func TestSite_GetChildren(t *testing.T) {
	site := Site{
		Url:       NewKSVString("https://example.com"),
		Name:      NewKSVString("Example"),
		AppId:     NewKSVString("app123"),
		GenericId: NewKSVString("gen123"),
		Notes:     NewKSVString("Some notes"),
		Entries:   KSVMap[*Entry]{},
	}

	children := site.GetChildren()
	assert.Contains(t, children, "url")
	assert.Contains(t, children, "name")
	assert.Contains(t, children, "appId")
	assert.Contains(t, children, "genericId")
	assert.Contains(t, children, "notes")
}

func TestSite_GetValues(t *testing.T) {
	site := Site{
		Url:       NewKSVString("https://example.com"),
		Name:      NewKSVString("Example"),
		AppId:     NewKSVString("app123"),
		GenericId: NewKSVString("gen123"),
		Notes:     NewKSVString("Some notes"),
		Entries:   KSVMap[*Entry]{},
	}

	values := site.GetValues()
	assert.Equal(t, "https://example.com", values["url"])
	assert.Equal(t, "Example", values["name"])
	assert.Equal(t, "app123", values["appId"])
	assert.Equal(t, "gen123", values["genericId"])
	assert.Equal(t, "Some notes", values["notes"])
}

func TestSite_String(t *testing.T) {
	site := Site{}
	assert.Equal(t, "", site.String())
}
