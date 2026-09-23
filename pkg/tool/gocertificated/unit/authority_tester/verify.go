package authority_tester

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
)

func Verify(
	root *authority.Authority,
	cluster *authority.Authority,
	leaf *x509.Certificate,
) error {
	roots := x509.NewCertPool()
	roots.AddCert(root.Material().Certificate)
	middle := x509.NewCertPool()
	middle.AddCert(cluster.Material().Certificate)
	o := x509.VerifyOptions{}
	o.Roots = roots
	o.Intermediates = middle
	o.KeyUsages = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	_, e := leaf.Verify(o)

	return e
}
