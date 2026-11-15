package upstream_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"net/http"
	"strconv"
	"testing"
)

func Answered(
	t *testing.T,
	status int,
) *http.Response {
	t.Helper()
	result := &http.Response{
		StatusCode: status,
		Status: key_value.Space(
			strconv.Itoa(status),
			http.StatusText(status),
		),
		Body: http.NoBody,
	}
	t.Cleanup(func() { errors.LogClose(result.Body) })

	return result
}
