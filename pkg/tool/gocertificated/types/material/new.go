package material

import (
	"crypto"
	"crypto/x509"
)

func New(
	certificate *x509.Certificate,
	key crypto.Signer,
) *Material {
	return &Material{Certificate: certificate, Key: key}
}
