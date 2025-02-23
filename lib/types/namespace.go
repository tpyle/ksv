package types

import "github.com/tpyle/ksv/lib/ksverrors"

type Namespace struct {
	Sites KSVMap[*Site] `json:"sites"`
}

func (n *Namespace) Validate() []error {
	var errs []error
	for _, s := range n.Sites {
		errs = append(errs, s.Validate()...)
	}
	return errs
}

func (n *Namespace) Get(path string) (Queryable, error) {
	if site, ok := n.Sites[path]; ok {
		return site, nil
	}
	return nil, ksverrors.ErrNoSuchPath
}

func (n *Namespace) Set(value string) error {
	return ksverrors.ErrCannotSet
}

func (n *Namespace) GetChildren() []string {
	children := make([]string, 0, len(n.Sites))
	for key := range n.Sites {
		children = append(children, key)
	}
	return children
}

func (n *Namespace) GetValues() map[string]string {
	values := make(map[string]string)
	for key, site := range n.Sites {
		for k, v := range site.GetValues() {
			values[key+"."+k] = v
		}
	}
	return values
}

func (n *Namespace) String() string {
	return "Namespace"
}
