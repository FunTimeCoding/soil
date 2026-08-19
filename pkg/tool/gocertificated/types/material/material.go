package material

import (
	"crypto"
	"crypto/x509"
)

type Material struct {
	Certificate *x509.Certificate
	Key         crypto.Signer
}
