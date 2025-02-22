package types

import (
	"testing"
)

func TestKSVString_Get(t *testing.T) {
	s := KSVString("test")
	result, err := s.Get("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != &s {
		t.Errorf("Expected %v, got %v", &s, result)
	}
}

func TestKSVString_Set(t *testing.T) {
	var s KSVString
	err := s.Set("test")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if s != KSVString("test") {
		t.Errorf("Expected test, got %v", s)
	}
}

func TestKSVString_Validate(t *testing.T) {
	s := KSVString("test")
	errs := s.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %v", errs)
	}
}

func TestKSVString_GetChildren(t *testing.T) {
	s := KSVString("test")
	children := s.GetChildren()
	if len(children) != 1 || children[0] != "" {
		t.Errorf("Expected [\"\"], got %v", children)
	}
}

func TestKSVString_GetValues(t *testing.T) {
	s := KSVString("test")
	values := s.GetValues()
	if len(values) != 1 || values[""] != "test" {
		t.Errorf("Expected {\"\": \"test\"}, got %v", values)
	}
}

func TestKSVString_String(t *testing.T) {
	s := KSVString("test")
	str := s.String()
	if str != "test" {
		t.Errorf("Expected test, got %v", str)
	}
}
