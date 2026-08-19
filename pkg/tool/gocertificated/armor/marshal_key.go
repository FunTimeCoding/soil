package armor

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func MarshalKey(k crypto.Signer) []byte {
	b, e := x509.MarshalPKCS8PrivateKey(k)
	errors.PanicOnError(e)

	return pem.EncodeToMemory(&pem.Block{Type: constant.KeyBlock, Bytes: b})
}
