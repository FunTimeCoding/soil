package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func NewDelete(locator string) *http.Request {
	result, e := http.NewRequest(constant.Delete, locator, nil)
	errors.PanicOnError(e)

	return result
}
