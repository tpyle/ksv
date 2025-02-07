package errors

import "fmt"

var (
	ErrGeneratorInvalidLength    = fmt.Errorf("invalid length")
	ErrGeneratorImpossibleParams = fmt.Errorf("impossible parameters")
)

var (
	ErrMissingIDPUrl   = fmt.Errorf("missing parameter: IDP URL is required")
	ErrMissingIDPName  = fmt.Errorf("missing parameter: IDP Name is required")
	ErrInvalidKnownIDP = fmt.Errorf("invalid KnownIDP")
)
