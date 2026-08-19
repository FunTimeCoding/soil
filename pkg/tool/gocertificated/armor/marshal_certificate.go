package armor

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func MarshalCertificate(c *x509.Certificate) []byte {
	return pem.EncodeToMemory(
		&pem.Block{Type: constant.CertificateBlock, Bytes: c.Raw},
	)
}
