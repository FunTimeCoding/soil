package jira

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func WrapError(
	r *jira.Response,
	e error,
) error {
	if e == nil {
		return nil
	}

	f, okay := errors.AsType[*jira.Error](e)
	detail := e.Error()

	if okay {
		parts := append([]string{}, f.ErrorMessages...)

		for field, message := range f.Errors {
			parts = append(parts, fmt.Sprintf("%s: %s", field, message))
		}

		if d := join.SemicolonSpace(parts); d != "" {
			detail = d
		}
	}

	if r != nil && r.Response != nil {
		if g := refusal.Status(r.StatusCode, detail); g != nil {
			return g
		}
	}

	if !okay {
		return e
	}

	status := ""

	if f.HTTPError != nil {
		status = f.HTTPError.Error()
	}

	body := notation.Marshal(response.NewError(f.ErrorMessages, f.Errors))

	return detail_error.New(detail, status).WithBody(body)
}
