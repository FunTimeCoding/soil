package signature

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
)

func SignControl(
	b []byte,
	k *rsa.PrivateKey,
) ([]byte, error) {
	h := sha1.Sum(b)

	return rsa.SignPKCS1v15(nil, k, crypto.SHA1, h[:])
}
