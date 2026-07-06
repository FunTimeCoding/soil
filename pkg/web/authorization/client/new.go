package client

import (
	"crypto/aes"
	"github.com/funtimecoding/soil/pkg/errors"
)

func New(
	issuer string,
	identifier string,
	secret string,
	signInPath string,
	callbackLocator string,
	encryptionKey []byte,
) *Client {
	block, e := aes.NewCipher(encryptionKey)
	errors.PanicOnError(e)
	seal, e := newSeal(block)
	errors.PanicOnError(e)

	return &Client{
		issuer:          issuer,
		identifier:      identifier,
		secret:          secret,
		signInPath:      signInPath,
		callbackLocator: callbackLocator,
		seal:            seal,
	}
}
