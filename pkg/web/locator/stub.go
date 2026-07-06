package locator

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
)

func Stub() *Locator {
	return &Locator{
		scheme:        constant.Secure,
		value:         url.Values{},
		fragmentValue: url.Values{},
	}
}
