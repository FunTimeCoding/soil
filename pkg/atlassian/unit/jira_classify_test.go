package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
	"testing"
)

func TestAnUnknownIssueIsNotFound(t *testing.T) {
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

func TestARejectedRequestIsValidation(t *testing.T) {
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

func TestAnUpstreamFailureKeepsItsDetail(t *testing.T) {
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

func TestAFailureWithoutResponseKeepsItsDetail(t *testing.T) {
	_, f := jiraFailure(http.StatusNotFound, []string{"alfa"}, nil)
	e := jira.WrapError(nil, f)
	assert.False(t, not_found.Is(e))
	_, okay := errors.AsType[*detail_error.Detail](e)
	assert.True(t, okay)
}
