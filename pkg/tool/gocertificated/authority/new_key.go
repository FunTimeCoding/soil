package authority

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/funtimecoding/soil/pkg/errors"
)

func newKey() crypto.Signer {
	k, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	errors.PanicOnError(e)

	return k
}
