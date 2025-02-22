package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/errors"
)

func TestEntry_Get(t *testing.T) {
	e := Entry{
		IsIDP:        KSVBool(true),
		IDP:          IDP{},
		Username:     KSVString("user"),
		Email:        KSVString("user@example.com"),
		Notes:        KSVString("some notes"),
		CustomFields: KSVMap{},
		SecretFields: KSVMap{},
	}

	// Test valid paths
	paths := []string{"isIDP", "idp", "username", "email", "notes"}
	for _, path := range paths {
		_, err := e.Get(path)
		assert.NoError(t, err)
	}

	// Test invalid path
	_, err := e.Get("invalid")
	assert.ErrorIs(t, err, errors.ErrNoSuchPath)
}

func TestEntry_Set(t *testing.T) {
	e := Entry{}
	err := e.Set("value")
	assert.ErrorIs(t, err, errors.ErrCannotSet)
}

func TestEntry_Validate(t *testing.T) {
	e := Entry{
		Username:     KSVString(""),
		IsIDP:        true,
		IDP:          IDP{},
		CustomFields: KSVMap{},
		SecretFields: KSVMap{},
	}
	errs := e.Validate()
	assert.NotEmpty(t, errs)
}

func TestEntry_GetChildren(t *testing.T) {
	e := Entry{
		IsIDP:        KSVBool(true),
		IDP:          IDP{},
		Username:     KSVString("user"),
		Email:        KSVString("user@example.com"),
		Notes:        KSVString("some notes"),
		CustomFields: KSVMap{},
		SecretFields: KSVMap{},
	}
	children := e.GetChildren()
	assert.NotEmpty(t, children)
}

func TestEntry_GetValues(t *testing.T) {
	e := Entry{
		IsIDP:        KSVBool(true),
		IDP:          IDP{},
		Username:     KSVString("user"),
		Email:        KSVString("user@example.com"),
		Notes:        KSVString("some notes"),
		CustomFields: KSVMap{},
		SecretFields: KSVMap{},
	}
	values := e.GetValues()
	assert.NotEmpty(t, values)
}

func TestEntry_String(t *testing.T) {
	e := Entry{}
	str := e.String()
	assert.Equal(t, "", str)
}
