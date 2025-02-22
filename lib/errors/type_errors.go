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
	ErrMissingUsername        = fmt.Errorf("missing parameter: Username is required")
	ErrMissingCustomFieldName = fmt.Errorf("missing parameter: Custom Field Name is required")
	ErrMissingSecretFieldName = fmt.Errorf("missing parameter: Secret Field Name is required")
)

var (
	ErrEmptyPathElement = fmt.Errorf("empty path element")
)

var (
	ErrMissingSiteUrlOrAppIdOrGenericId = fmt.Errorf("missing parameter: site must have URL, AppId, or GenericId")
	ErrMissingSiteName                  = fmt.Errorf("missing parameter: Site Name is required")
)

var (
	ErrEmptyLocalStorage           = fmt.Errorf("local storage is empty")
	ErrUnsupportedLocalStorageType = fmt.Errorf("unsupported local storage type")
)

var (
	ErrInvalidPath = fmt.Errorf("invalid path")
	ErrNoSuchPath  = fmt.Errorf("no such path")
	ErrCannotSet   = fmt.Errorf("cannot set this value")
)

var (
	ErrClipboardUnavailable = fmt.Errorf("clipboard unavailable")
)
