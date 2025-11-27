package types

import (
	"io"

	"github.com/tpyle/ksv/lib/enc"
	"github.com/tpyle/ksv/lib/ksverrors"
)

type Namespace struct {
	Sites KSVMap[*Site] `json:"sites"`
	IDPs  KSVMap[*IDP]  `json:"idps"`
}

func DecryptNamespace(input io.Reader, key []byte) (*Namespace, error) {
	var ns Namespace
	err := enc.DecryptAndUnmarshal(input, key, &ns)
	if err != nil {
		return nil, err
	}
	return &ns, nil
}

func (n *Namespace) Encrypt(output io.Writer, key []byte) error {
	return enc.MarshalAndEncrypt(n, key, output)
}

func NewNamespace() Namespace {
	return Namespace{
		Sites: make(KSVMap[*Site]),
		IDPs:  make(KSVMap[*IDP]),
	}
}

func (n *Namespace) AddSite(site Site) error {
	if n.Sites == nil {
		n.Sites = make(KSVMap[*Site])
	}
	if _, ok := n.Sites[site.Name.Value]; ok {
		return ksverrors.ErrSiteAlreadyExists
	}
	n.Sites[site.Name.Value] = &site
	return nil
}

func (n *Namespace) AddIDP(idp IDP) {
	if n.IDPs == nil {
		n.IDPs = make(KSVMap[*IDP])
	}
	n.IDPs[idp.Name.Value] = &idp
}

func (n *Namespace) GetIDPByName(name string) (*IDP, bool) {
	idp, ok := n.IDPs[name]
	return idp, ok
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
