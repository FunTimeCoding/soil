package gitlab

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"net/http"
)

func wrapError(e error) error {
	if e == nil {
		return nil
	}

	f, okay := errors.AsType[*gitlab.ErrorResponse](e)

	if !okay {
		if g := connection.Classify(e); g != nil {
			return g
		}

		return e
	}

	detail := f.Message

	if detail == "" {
		detail = f.Error()
	}

	if f.StatusCode == http.StatusPreconditionFailed {
		return conflict.Format("%s", detail)
	}

	if g := refusal.NewTable().
		With(http.StatusForbidden, errorsConstant.ForbiddenClass).
		With(http.StatusUnprocessableEntity, errorsConstant.ValidationClass).
		Classify(f.StatusCode, detail); g != nil {
		return g
	}
	status := fmt.Sprintf("%d %s", f.StatusCode, http.StatusText(f.StatusCode))

	return detail_error.New(detail, status).WithBody(f.Body)
}
