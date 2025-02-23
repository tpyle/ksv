package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/ksverrors"
)

func TestKSVMap_Get(t *testing.T) {
	val := KSVString{
		Value: "value1",
	}
	m := KSVMap[*KSVString]{"key1": &val}

	t.Run("existing key", func(t *testing.T) {
		val, err := m.Get("key1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", val.String())
	})

	t.Run("non-existing key", func(t *testing.T) {
		_, err := m.Get("")
		assert.ErrorIs(t, err, ksverrors.ErrNoSuchPath)
	})

	t.Run("empty path", func(t *testing.T) {
		_, err := m.Get("non-existing")
		assert.ErrorIs(t, err, ksverrors.ErrNoSuchPath)
	})
}

func TestKSVMap_Set(t *testing.T) {
	m := KSVMap[*KSVString]{}
	err := m.Set("value")
	assert.ErrorIs(t, err, ksverrors.ErrCannotSet)
}

func TestKSVMap_Validate(t *testing.T) {
	m := KSVMap[*KSVString]{}
	errs := m.Validate()
	assert.Empty(t, errs)
}

func TestKSVMap_GetChildren(t *testing.T) {
	val := KSVString{
		Value: "value1",
	}
	m := KSVMap[*KSVString]{"key1": &val}
	children := m.GetChildren()
	assert.Contains(t, children, "key1")
}

func TestKSVMap_GetValues(t *testing.T) {
	val := KSVString{
		Value: "value1",
	}
	m := KSVMap[*KSVString]{"key1": &val}
	values := m.GetValues()
	assert.Equal(t, "value1", values["key1"])
}

func TestKSVMap_String(t *testing.T) {
	m := KSVMap[*KSVString]{}
	assert.Equal(t, "", m.String())
}
