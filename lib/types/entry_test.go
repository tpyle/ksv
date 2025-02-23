package types

// func TestEntry_Get(t *testing.T) {
// 	e := Entry{
// 		IsIDP:        KSVBoolTrue,
// 		IDP:          IDP{},
// 		Username:     NewKSVString("user"),
// 		Email:        NewKSVString("user@example.com"),
// 		Notes:        NewKSVString("some notes"),
// 		CustomFields: KSVMap[KSVString]{},
// 		SecretFields: KSVMap[KSVString]{},
// 	}

// 	// Test valid paths
// 	paths := []string{"isIDP", "idp", "username", "email", "notes"}
// 	for _, path := range paths {
// 		_, err := e.Get(path)
// 		assert.NoError(t, err)
// 	}

// 	// Test invalid path
// 	_, err := e.Get("invalid")
// 	assert.ErrorIs(t, err, ksverrors.ErrNoSuchPath)
// }

// func TestEntry_Set(t *testing.T) {
// 	e := Entry{}
// 	err := e.Set("value")
// 	assert.ErrorIs(t, err, ksverrors.ErrCannotSet)
// }

// func TestEntry_Validate(t *testing.T) {
// 	e := Entry{
// 		Username:     KSVStringEmpty,
// 		IsIDP:        KSVBoolTrue,
// 		IDP:          IDP{},
// 		CustomFields: KSVMap[KSVString]{},
// 		SecretFields: KSVMap[KSVString]{},
// 	}
// 	errs := e.Validate()
// 	assert.NotEmpty(t, errs)
// }

// func TestEntry_GetChildren(t *testing.T) {
// 	e := Entry{
// 		IsIDP:        KSVBoolTrue,
// 		IDP:          IDP{},
// 		Username:     NewKSVString("user"),
// 		Email:        NewKSVString("user@example.com"),
// 		Notes:        NewKSVString("some notes"),
// 		CustomFields: KSVMap[KSVString]{},
// 		SecretFields: KSVMap[KSVString]{},
// 	}
// 	children := e.GetChildren()
// 	assert.NotEmpty(t, children)
// }

// func TestEntry_GetValues(t *testing.T) {
// 	e := Entry{
// 		IsIDP:        KSVBoolTrue,
// 		IDP:          IDP{},
// 		Username:     NewKSVString("user"),
// 		Email:        NewKSVString("user@example.com"),
// 		Notes:        NewKSVString("some notes"),
// 		CustomFields: KSVMap[KSVString]{},
// 		SecretFields: KSVMap[KSVString]{},
// 	}
// 	values := e.GetValues()
// 	assert.NotEmpty(t, values)
// }

// func TestEntry_String(t *testing.T) {
// 	e := Entry{}
// 	str := e.String()
// 	assert.Equal(t, "", str)
// }
