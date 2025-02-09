package types

import (
	"testing"
)

func TestEntry_Validate(t *testing.T) {
	validEntry := Entry{
		DistinguishingName: "validName",
		IsIDP:              false,
		Username:           "validUser",
		Email:              "valid@example.com",
	}

	invalidEntry := Entry{
		DistinguishingName: "",
		IsIDP:              false,
		Username:           "",
		Email:              "",
	}

	errs := invalidEntry.Validate()
	if len(errs) == 0 {
		t.Errorf("Expected validation errors, got none")
	}

	errs = validEntry.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no validation errors, got %d", len(errs))
	}
}
