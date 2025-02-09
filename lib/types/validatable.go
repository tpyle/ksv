package types

type Validatable interface {
	Validate() []error
}
