package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func NewGet(
	format string,
	a ...any,
) *http.Request {
	if len(a) > 0 {
		format = fmt.Sprintf(format, a...)
	}

	result, e := http.NewRequest(constant.Get, format, nil)
	errors.PanicOnError(e)

	return result
}
