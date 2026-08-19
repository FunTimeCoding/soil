package authority

import (
	"crypto"
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/policy"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
)

func (a *Authority) Sign(
	r *issue_request.Request,
	public crypto.PublicKey,
) *x509.Certificate {
	p := policy.New(r.Kind)
	p.Inherit(r.Name, distinguished_name.From(a.material.Certificate.Subject))
	p.Validate(r.Name)

	return newCertificate(
		newTemplate(r),
		a.material.Certificate,
		public,
		a.material.Key,
	)
}
