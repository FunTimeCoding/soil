package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
	"net/url"
	"testing"
)

func TestUnknownIssueIsNotFound(t *testing.T) {
	e := jira.WrapError(
		jiraFailure(
			http.StatusNotFound,
			[]string{
				"Issue does not exist or you do not have permission to see it.",
			},
			nil,
		),
	)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "Issue does not exist", e.Error())
}

func TestRejectedRequestIsValidation(t *testing.T) {
	e := jira.WrapError(
		jiraFailure(
			http.StatusBadRequest,
			nil,
			map[string]string{"comment": "Comment body can not be empty!"},
		),
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "Comment body can not be empty!", e.Error())
}

func TestUpstreamFailureKeepsItsDetail(t *testing.T) {
	e := jira.WrapError(
		jiraFailure(
			http.StatusInternalServerError,
			[]string{"Internal server error"},
			nil,
		),
	)
	assert.False(t, not_found.Is(e))
	assert.False(t, validation.Is(e))
	_, okay := errors.AsType[*detail_error.Detail](e)
	assert.True(t, okay)
	assert.StringContains(t, "Internal server error", e.Error())
}

func TestLostConnectionWithoutResponseIsDropped(t *testing.T) {
	e := jira.WrapError(
		nil,
		fmt.Errorf(
			"No response returned: %w",
			&url.Error{
				Op:  "Get",
				URL: "https://jira.example.org/rest/api/2/search/jql",
				Err: errors.New("http2: client connection lost"),
			},
		),
	)
	f, okay := errors.AsType[*connection.Failure](e)
	assert.True(t, okay)
	assert.True(t, dropped.Is(e))
	assert.String(t, "jira.example.org", f.Host)
}

func TestFailureWithoutResponseKeepsItsDetail(t *testing.T) {
	_, f := jiraFailure(http.StatusNotFound, []string{"alfa"}, nil)
	e := jira.WrapError(nil, f)
	assert.False(t, not_found.Is(e))
	_, okay := errors.AsType[*detail_error.Detail](e)
	assert.True(t, okay)
}
