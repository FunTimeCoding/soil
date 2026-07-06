package base64

import (
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Decode(s string) string {
	result, e := base64.StdEncoding.DecodeString(s)
	errors.PanicOnError(e)

	return string(result)
}
