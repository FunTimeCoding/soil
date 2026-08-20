package server

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func jiraMessage(e *jira.Error) string {
	parts := append([]string{}, e.ErrorMessages...)

	for field, message := range e.Errors {
		parts = append(parts, fmt.Sprintf("%s: %s", field, message))
	}

	return join.SemicolonSpace(parts)
}
