package types

import (
	"testing"
)

func TestNamespace_Validate(t *testing.T) {
	validSite := Site{
		Url:  "https://example.com",
		Name: "Example",
	}
	invalidSite := Site{}

	namespace := Namespace{
		Sites: []Site{
			validSite,
			invalidSite,
		},
	}

	errs := namespace.Validate()
	if len(errs) == 0 {
		t.Errorf("Expected validation errors, got none")
	}
}
