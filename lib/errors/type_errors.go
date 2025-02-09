package errors

import "fmt"

var (
	ErrGeneratorInvalidLength    = fmt.Errorf("invalid length")
	ErrGeneratorImpossibleParams = fmt.Errorf("impossible parameters")
	ErrUnsupportedGenerator      = fmt.Errorf("unsupported generator")
)

var (
	ErrMissingIDPUrl   = fmt.Errorf("missing parameter: IDP URL is required")
	ErrMissingIDPName  = fmt.Errorf("missing parameter: IDP Name is required")
	ErrInvalidKnownIDP = fmt.Errorf("invalid KnownIDP")
)

var (
	ErrMissingDistinguishingName = fmt.Errorf("missing parameter: Distinguishing Name is required")
	ErrMissingCustomFieldName    = fmt.Errorf("missing parameter: Custom Field Name is required")
	ErrMissingSecretFieldName    = fmt.Errorf("missing parameter: Secret Field Name is required")
)

var (
	ErrMissingSiteUrlOrAppIdOrGenericId = fmt.Errorf("missing parameter: site must have URL, AppId, or GenericId")
	ErrMissingSiteName                  = fmt.Errorf("missing parameter: Site Name is required")
)

var (
	ErrUnsupportedLocalStorageType = fmt.Errorf("unsupported local storage type")
)
