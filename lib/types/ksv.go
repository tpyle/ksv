package types

type KSV struct {
	Sites []Site `json:"sites"`
}

func (k *KSV) Validate() []error {
	var errs []error
	for _, s := range k.Sites {
		errs = append(errs, s.Validate()...)
	}
	return errs
}
