package armor

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func DecodeCertificate(b []byte) *x509.Certificate {
	c, e := x509.ParseCertificate(decodeBlock(b, constant.CertificateBlock))
	errors.PanicOnError(e)

	return c
}
