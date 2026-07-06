package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func DecodeLocator(s string) string {
	result, e := url.QueryUnescape(s)
	errors.PanicOnError(e)

	return result
}
