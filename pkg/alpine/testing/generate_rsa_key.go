package testing

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/errors"
)

func GenerateRSAKey(bits int) *rsa.PrivateKey {
	result, e := rsa.GenerateKey(rand.Reader, bits)
	errors.PanicOnError(e)

	return result
}
