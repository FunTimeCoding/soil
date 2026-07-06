package web

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func NewPut(
	locator string,
	body string,
) *http.Request {
	result, e := http.NewRequest(
		constant.Put,
		locator,
		bytes.NewBuffer([]byte(body)),
	)
	errors.PanicOnError(e)

	return result
}
