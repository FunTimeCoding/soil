package name_constraint

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func (c *Constraint) Apply(t *x509.Certificate) {
	t.PermittedDNSDomains = c.PermittedDomain
	t.PermittedDNSDomainsCritical = true
	t.PermittedIPRanges = c.PermittedAddress
	t.ExcludedEmailAddresses = []string{constant.AnyName}
	t.ExcludedURIDomains = []string{constant.AnyName}
}
