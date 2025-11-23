package ksverrors

import (
	"errors"
)

var (
	ErrCliNoMoreArgs           = errors.New("no additional arguments are allowed")
	ErrCliFailedToGetKeystore  = errors.New("failed to get keystore from context")
	ErrCliFailedToSaveKeystore = errors.New("failed to save keystore to context")
)

var (
	ErrCliSiteDoesNotExist = errors.New("site does not exist")
)

var (
	ErrCliCouldNotEncodeYaml = errors.New("could not encode YAML")
	ErrCliCouldNotEncodeJSON = errors.New("could not encode JSON")
	ErrCliCouldNotWrite      = errors.New("could not write output")
)
