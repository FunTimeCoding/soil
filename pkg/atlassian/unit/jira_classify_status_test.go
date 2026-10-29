package unit

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
	"testing"
)

func TestARefusedProjectIsForbidden(t *testing.T) {
	e := atlassian.WrapError(
		jiraFailure(
			http.StatusForbidden,
			[]string{"You do not have permission to see this project."},
			nil,
		),
	)
	assert.True(t, forbidden.Is(e))
	assert.StringContains(t, "permission to see this project", e.Error())
}

func TestAnUnavailableJiraIsUnavailable(t *testing.T) {
	e := atlassian.WrapError(
		jiraFailure(
			http.StatusServiceUnavailable,
			[]string{"Down for upgrade"},
			nil,
		),
	)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "Down for upgrade", e.Error())
}

func TestAnUnavailableJiraWithoutNotationIsUnavailable(t *testing.T) {
	e := atlassian.WrapError(
		&jira.Response{
			Response: &http.Response{StatusCode: http.StatusServiceUnavailable},
		},
		errors.New("503 Service Unavailable: <html>maintenance</html>"),
	)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "maintenance", e.Error())
}

func TestARejectedTokenStaysAFault(t *testing.T) {
	e := atlassian.WrapError(
		jiraFailure(
			http.StatusUnauthorized,
			[]string{"You are not authenticated."},
			nil,
		),
	)
	assert.False(t, forbidden.Is(e))
	_, okay := errors.AsType[*detail_error.Detail](e)
	assert.True(t, okay)
}
