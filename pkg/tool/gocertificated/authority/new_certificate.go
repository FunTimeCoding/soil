package authority

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
)

func newCertificate(
	template *x509.Certificate,
	issuer *x509.Certificate,
	public crypto.PublicKey,
	signer crypto.Signer,
) *x509.Certificate {
	b, e := x509.CreateCertificate(
		rand.Reader,
		template,
		issuer,
		public,
		signer,
	)
	errors.PanicOnError(e)
	c, f := x509.ParseCertificate(b)
	errors.PanicOnError(f)

	return c
}
