package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
	"strings"
)

func ParseLocator(s string) *url.URL {
	if !strings.HasPrefix(s, constant.SecurePrefix) &&
		!strings.HasPrefix(s, constant.InsecurePrefix) {
		panic("locator must have a scheme")
	}

	result, e := url.Parse(s)
	errors.PanicOnError(e)

	return result
}
