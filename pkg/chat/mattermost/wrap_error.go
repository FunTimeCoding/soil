package mattermost

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func wrapError(e error) error {
	f, okay := errors.AsType[*model.AppError](e)

	if !okay {
		return e
	}

	status := fmt.Sprintf("%d %s", f.StatusCode, http.StatusText(f.StatusCode))

	return detail_error.New(f.Message, status).WithBody(notation.Marshal(f))
}
