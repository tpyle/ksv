package types

import (
	"testing"
)

func TestKSV_Validate(t *testing.T) {
	validSite := Site{ /* initialize with valid data */ }
	invalidSite := Site{ /* initialize with invalid data */ }

	ksv := KSV{
		Sites: []Site{validSite, invalidSite},
	}

	errs := ksv.Validate()
	if len(errs) == 0 {
		t.Errorf("Expected validation errors, got none")
	}
}
