package types

type Queryable interface {
	Get(pathElement string) (Queryable, error)
	Set(value string) error
	Validate() []error
	GetChildren() []string
	GetValues() map[string]string
	String() string
}
