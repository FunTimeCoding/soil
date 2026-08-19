package policy

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func New(kind constant.CertificateKind) *Policy {
	p := &Policy{}

	switch kind {
	case constant.KindRoot:
		p.RequireCountry = true
		p.RequireProvince = true
		p.RequireOrganization = true
	case constant.KindIntermediate:
		p.RequireOrganization = true
		p.InheritIssuer = true
	}

	return p
}
