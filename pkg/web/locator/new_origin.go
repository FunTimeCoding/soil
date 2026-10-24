package locator

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
	"strconv"
)

func NewOrigin(origin string) *Locator {
	u, e := url.Parse(origin)

	if e != nil {
		return New(origin)
	}

	result := New(u.Hostname())

	if u.Scheme == constant.Insecure {
		result.Insecure()
	}

	if port, f := strconv.Atoi(u.Port()); f == nil {
		result.Port(port)
	}

	return result
}
