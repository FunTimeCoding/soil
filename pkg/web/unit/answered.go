package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"testing"
)

func answered(
	t *testing.T,
	status int,
) *http.Response {
	t.Helper()
	result := &http.Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Body:       http.NoBody,
	}
	t.Cleanup(func() { errors.LogClose(result.Body) })

	return result
}
