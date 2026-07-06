package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func NewPropfind(locator string) *http.Request {
	result, e := http.NewRequest(constant.Propfind, locator, nil)
	errors.PanicOnError(e)

	return result
}
