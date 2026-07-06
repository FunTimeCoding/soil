package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func NewPutBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		constant.Put,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
