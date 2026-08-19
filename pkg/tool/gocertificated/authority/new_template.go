package authority

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"net"
	"time"
)

func newTemplate(r *issue_request.Request) *x509.Certificate {
	now := time.Now()
	t := &x509.Certificate{}
	t.SerialNumber = newSerial()
	t.Subject = r.Name.Subject()
	t.NotBefore = now.Add(-constant.ClockSkew)
	t.NotAfter = now.AddDate(r.ValidYear, 0, r.ValidDay)
	t.BasicConstraintsValid = true

	switch r.Kind {
	case constant.KindRoot:
		t.IsCA = true
		t.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	case constant.KindIntermediate:
		t.IsCA = true
		t.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
		t.MaxPathLen = 0
		t.MaxPathLenZero = true
	case constant.KindServer:
		t.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
		t.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	case constant.KindClient:
		t.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
		t.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
	}

	for _, h := range r.Host {
		if a := net.ParseIP(h); a != nil {
			t.IPAddresses = append(t.IPAddresses, a)

			continue
		}

		t.DNSNames = append(t.DNSNames, h)
	}

	if r.Constraint != nil {
		r.Constraint.Apply(t)
	}

	return t
}
