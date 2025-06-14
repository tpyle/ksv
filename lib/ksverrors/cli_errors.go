package ksverrors

import "fmt"

var (
	ErrCliNoMoreArgs           = fmt.Errorf("no additional arguments are allowed")
	ErrCliFailedToGetKeystore  = fmt.Errorf("failed to get keystore from context")
	ErrCliFailedToSaveKeystore = fmt.Errorf("failed to save keystore to context")
)

var (
	ErrCliSiteDoesNotExist = fmt.Errorf("site does not exist")
)

var (
	ErrCliCouldNotEncodeYaml = fmt.Errorf("could not encode YAML")
	ErrCliCouldNotEncodeJSON = fmt.Errorf("could not encode JSON")
	ErrCliCouldNotWrite      = fmt.Errorf("could not write output")
)
