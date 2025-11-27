package ksverrors

import "errors"

var (
	ErrGeneratorInvalidLength    = errors.New("invalid length")
	ErrGeneratorImpossibleParams = errors.New("impossible parameters")
	ErrUnsupportedGenerator      = errors.New("unsupported generator")
)

var (
	ErrMissingIDPUrl   = errors.New("missing parameter: IDP URL is required")
	ErrMissingIDPName  = errors.New("missing parameter: IDP Name is required")
	ErrInvalidKnownIDP = errors.New("invalid KnownIDP")
)

var (
	ErrMissingUsername        = errors.New("missing parameter: Username is required")
	ErrMissingCustomFieldName = errors.New("missing parameter: Custom Field Name is required")
	ErrMissingSecretFieldName = errors.New("missing parameter: Secret Field Name is required")
	ErrEmptyValue             = errors.New("empty value")
)

var (
	ErrEmptyPathElement = errors.New("empty path element")
)

var (
	ErrMissingSiteUrlOrAppIdOrGenericId = errors.New("missing parameter: site must have URL, AppId, or GenericId")
	ErrMissingSiteName                  = errors.New("missing parameter: Site Name is required")
)

var (
	ErrEmptyLocalStorage           = errors.New("local storage is empty")
	ErrUnsupportedLocalStorageType = errors.New("unsupported local storage type")
)

var (
	ErrInvalidPath = errors.New("invalid path")
	ErrNoSuchPath  = errors.New("no such path")
	ErrCannotSet   = errors.New("cannot set this value")
)

var (
	ErrClipboardUnavailable = errors.New("clipboard unavailable")
)

var (
	ErrEntryAlreadyExists = errors.New("entry already exists")
	ErrEntryDoesNotExist  = errors.New("entry does not exist")
	ErrSiteAlreadyExists  = errors.New("site already exists")
	ErrSiteDoesNotExist   = errors.New("site does not exist")
	ErrIDPAlreadyExists   = errors.New("IDP already exists")

	ErrNoSuchNamespace = errors.New("no such namespace")
)
