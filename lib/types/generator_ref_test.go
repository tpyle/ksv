package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/errors"
)

func TestGeneratorReference_UnmarshalJSON(t *testing.T) {
	jsonStr := `{"ref": "testRef", "params": {"key": "value"}}`
	var gr GeneratorReference
	err := json.Unmarshal([]byte(jsonStr), &gr)
	assert.NoError(t, err)
	assert.Equal(t, "testRef", gr.Ref.String())
	val, err := gr.Params.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", val.String())
	assert.False(t, gr.Supported)
	assert.NotNil(t, gr.Generator)
}

func TestGeneratorReference_Get(t *testing.T) {
	ref := KSVString("testRef")
	pval := KSVString("value")
	gr := &GeneratorReference{
		Ref:    ref,
		Params: KSVMap{"key": &pval},
	}

	val, err := gr.Get("ref")
	assert.NoError(t, err)
	assert.Equal(t, "testRef", val.String())

	val, err = gr.Get("params")
	assert.NoError(t, err)
	subval, err := val.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", subval.String())

	_, err = gr.Get("nonexistent")
	assert.Error(t, err)
	assert.Equal(t, errors.ErrNoSuchPath, err)
}

func TestGeneratorReference_Set(t *testing.T) {
	gr := &GeneratorReference{}
	err := gr.Set("newValue")
	assert.Error(t, err)
	assert.Equal(t, errors.ErrCannotSet, err)
}

func TestGeneratorReference_Validate(t *testing.T) {
	gr := &GeneratorReference{}
	errs := gr.Validate()
	assert.Empty(t, errs)
}

func TestGeneratorReference_GetChildren(t *testing.T) {
	pval := KSVString("value")
	gr := &GeneratorReference{
		Ref: KSVString("testRef"),
		Params: KSVMap{
			"key": &pval,
		},
	}

	children := gr.GetChildren()
	assert.Contains(t, children, "ref")
	assert.Contains(t, children, "params/key")
}

func TestGeneratorReference_GetValues(t *testing.T) {
	pval := KSVString("value")
	gr := &GeneratorReference{
		Ref: KSVString("testRef"),
		Params: KSVMap{
			"key": &pval,
		},
	}

	values := gr.GetValues()
	assert.Equal(t, "testRef", values["ref"])
	assert.Equal(t, "value", values["params/key"])
}
