package types

import (
	"testing"
)

func TestKSVBool_Get(t *testing.T) {
	b := KSVBool(true)
	result, err := b.Get("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != &b {
		t.Errorf("Expected %v, got %v", &b, result)
	}
}

func TestKSVBool_Set(t *testing.T) {
	b := KSVBool(false)
	err := b.Set("true")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if b != KSVBool(true) {
		t.Errorf("Expected true, got %v", b)
	}
}

func TestKSVBool_Validate(t *testing.T) {
	b := KSVBool(true)
	errs := b.Validate()
	if len(errs) != 0 {
		t.Errorf("Expected no errors, got %v", errs)
	}
}

func TestKSVBool_GetChildren(t *testing.T) {
	b := KSVBool(true)
	children := b.GetChildren()
	if len(children) != 1 || children[0] != "" {
		t.Errorf("Expected [\"\"], got %v", children)
	}
}

func TestKSVBool_GetValues(t *testing.T) {
	b := KSVBool(true)
	values := b.GetValues()
	if len(values) != 1 || values[""] != "true" {
		t.Errorf("Expected {\"\": \"true\"}, got %v", values)
	}
}

func TestKSVBool_String(t *testing.T) {
	b := KSVBool(true)
	str := b.String()
	if str != "true" {
		t.Errorf("Expected true, got %v", str)
	}
}
