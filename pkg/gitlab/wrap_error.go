package gitlab

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"net/http"
)

func wrapError(e error) error {
	f, okay := errors.AsType[*gitlab.ErrorResponse](e)

	if !okay {
		return e
	}

	detail := f.Message

	if detail == "" {
		detail = f.Error()
	}

	status := fmt.Sprintf("%d %s", f.StatusCode, http.StatusText(f.StatusCode))

	return detail_error.New(detail, status).WithBody(f.Body)
}
