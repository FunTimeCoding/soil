package armor

import (
	"crypto"
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func DecodeKey(b []byte) crypto.Signer {
	k, e := x509.ParsePKCS8PrivateKey(decodeBlock(b, constant.KeyBlock))
	errors.PanicOnError(e)
	s, okay := k.(crypto.Signer)

	if !okay {
		errors.PanicOnError(errors.Format("key cannot sign", k))
	}

	return s
}
