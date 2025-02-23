package types

// func TestNamespace_Validate(t *testing.T) {
// 	validSite := Site{
// 		Url:  "https://example.com",
// 		Name: "Example",
// 	}
// 	invalidSite := Site{}
// 	namespace := Namespace{
// 		Sites: KSVMap{
// 			"validSite":   &validSite,
// 			"invalidSite": &invalidSite,
// 		},
// 	}
// 	errs := namespace.Validate()
// 	if len(errs) == 0 {
// 		t.Errorf("Expected validation errors, got none")
// 	}
// }

// func TestNamespace_Get(t *testing.T) {
// 	namespace := Namespace{
// 		Sites: KSVMap{
// 			"example": &Site{Url: "https://example.com", Name: "Example"},
// 		},
// 	}

// 	_, err := namespace.Get("example")
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	_, err = namespace.Get("nonexistent")
// 	if err == nil {
// 		t.Errorf("Expected error, got none")
// 	}
// }

// func TestNamespace_Set(t *testing.T) {
// 	namespace := Namespace{}

// 	err := namespace.Set("value")
// 	if err != ksverrors.ErrCannotSet {
// 		t.Errorf("Expected ErrCannotSet, got %v", err)
// 	}
// }

// func TestNamespace_GetChildren(t *testing.T) {
// 	namespace := Namespace{
// 		Sites: KSVMap{
// 			"example": &Site{Url: "https://example.com", Name: "Example"},
// 		},
// 	}

// 	children := namespace.GetChildren()
// 	if len(children) == 0 {
// 		t.Errorf("Expected children, got none")
// 	}
// }

// func TestNamespace_GetValues(t *testing.T) {
// 	namespace := Namespace{
// 		Sites: KSVMap{
// 			"example": &Site{Url: "https://example.com", Name: "Example"},
// 		},
// 	}

// 	values := namespace.GetValues()
// 	if len(values) == 0 {
// 		t.Errorf("Expected values, got none")
// 	}
// }

// func TestNamespace_String(t *testing.T) {
// 	namespace := Namespace{}

// 	str := namespace.String()
// 	if str == "" {
// 		t.Errorf("Expected non-empty string, got empty string")
// 	}
// }
