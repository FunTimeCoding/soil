package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func getWithContext(
	x context.Context,
	locator string,
) *http.Request {
	result, e := http.NewRequestWithContext(x, http.MethodGet, locator, nil)
	errors.PanicOnError(e)

	return result
}
