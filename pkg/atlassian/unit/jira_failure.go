package unit

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	"net/http"
)

func jiraFailure(
	status int,
	messages []string,
	fields map[string]string,
) (*jira.Response, error) {
	return &jira.Response{Response: &http.Response{StatusCode: status}},
		&jira.Error{
			HTTPError:     errors.New("request failed"),
			ErrorMessages: messages,
			Errors:        fields,
		}
}
