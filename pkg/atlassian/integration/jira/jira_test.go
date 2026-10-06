//go:build local

package jira

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestSearchReadsThroughTheRequester(t *testing.T) {
	result, e := jira.NewEnvironment().SearchLimitV3(
		1,
		"project = %s ORDER BY created DESC",
		environment.Required(constant.JiraDefaultProjectKeyEnvironment),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
}

func TestUnboundedSearchIsRefused(t *testing.T) {
	_, e := jira.NewEnvironment().SearchLimitV3(1, "ORDER BY created DESC")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "Unbounded JQL queries are not allowed", e.Error())
}

func TestLibraryClientsStillAnswer(t *testing.T) {
	u, e := jira.NewEnvironment().User()
	assert.FatalOnError(t, e)
	assert.True(t, u != nil)
}
