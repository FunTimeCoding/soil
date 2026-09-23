package authority_tester

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func NewSigningRequest(host string) string {
	k, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	errors.PanicOnError(e)
	t := &x509.CertificateRequest{}
	t.Subject = pkix.Name{CommonName: host}
	t.DNSNames = []string{host}
	b, f := x509.CreateCertificateRequest(rand.Reader, t, k)
	errors.PanicOnError(f)

	return string(
		pem.EncodeToMemory(
			&pem.Block{Type: constant.SigningRequestBlock, Bytes: b},
		),
	)
}
