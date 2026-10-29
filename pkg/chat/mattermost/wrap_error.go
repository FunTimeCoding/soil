package mattermost

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func wrapError(
	r *model.Response,
	e error,
) error {
	if e == nil || r == nil || r.StatusCode == 0 {
		return e
	}

	f, okay := errors.AsType[*model.AppError](e)
	reason := e.Error()

	if okay {
		reason = f.Message
	}

	if g := refusal.Status(r.StatusCode, reason); g != nil {
		return g
	}

	status := fmt.Sprintf("%d %s", r.StatusCode, http.StatusText(r.StatusCode))

	if !okay {
		return detail_error.New(reason, status)
	}

	return detail_error.New(f.Message, status).WithBody(notation.Marshal(f))
}
