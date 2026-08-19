package policy

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"

func (p *Policy) Inherit(
	subject *distinguished_name.Name,
	issuer *distinguished_name.Name,
) {
	if !p.InheritIssuer {
		return
	}

	subject.Country = issuer.Country
	subject.Province = issuer.Province
	subject.Organization = issuer.Organization
}
