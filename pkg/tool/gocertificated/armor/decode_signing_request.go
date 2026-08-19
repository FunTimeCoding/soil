package armor

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func DecodeSigningRequest(b []byte) *x509.CertificateRequest {
	r, e := x509.ParseCertificateRequest(
		decodeBlock(b, constant.SigningRequestBlock),
	)
	errors.PanicOnError(e)
	errors.PanicOnError(r.CheckSignature())

	return r
}
