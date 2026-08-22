package jira

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func wrapError(e error) error {
	f, okay := errors.AsType[*jira.Error](e)

	if !okay {
		return e
	}

	parts := append([]string{}, f.ErrorMessages...)

	for field, message := range f.Errors {
		parts = append(parts, fmt.Sprintf("%s: %s", field, message))
	}

	detail := join.SemicolonSpace(parts)

	if detail == "" {
		detail = f.Error()
	}

	status := ""

	if f.HTTPError != nil {
		status = f.HTTPError.Error()
	}

	body := notation.Marshal(ErrorPayload{f.ErrorMessages, f.Errors})

	return detail_error.New(detail, status).WithBody(body)
}
