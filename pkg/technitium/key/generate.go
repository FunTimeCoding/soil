package key

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
)

func Generate(
	name string,
	algorithm string,
) *Key {
	b := make([]byte, constant.SecretSize)
	_, e := rand.Read(b)
	errors.PanicOnError(e)

	return &Key{
		Name:      name,
		Secret:    base64.StdEncoding.EncodeToString(b),
		Algorithm: algorithm,
	}
}
