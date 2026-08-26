package package_server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"github.com/funtimecoding/soil/pkg/errors"
)

func SignControl(
	b []byte,
	k *rsa.PrivateKey,
) []byte {
	h := sha1.Sum(b)
	result, e := rsa.SignPKCS1v15(nil, k, crypto.SHA1, h[:])
	errors.PanicOnError(e)

	return result
}
