package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func NewPatchBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		constant.Patch,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
