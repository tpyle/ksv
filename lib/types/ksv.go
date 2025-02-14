package types

type KSV struct {
	Namespaces map[string]Namespace `json:"namespaces"`

	// Default Namespace
	DefaultNamespace Namespace `json:"default_namespace"`
}

func (k *KSV) Validate() []error {
	var errs []error
	errs = append(errs, k.DefaultNamespace.Validate()...)

	for _, n := range k.Namespaces {
		errs = append(errs, n.Validate()...)
	}

	return errs
}
