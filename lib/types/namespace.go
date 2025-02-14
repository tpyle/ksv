package types

type Namespace struct {
	Sites []Site `json:"sites"`
}

func (n *Namespace) Validate() []error {
	var errs []error
	for _, s := range n.Sites {
		errs = append(errs, s.Validate()...)
	}
	return errs
}
