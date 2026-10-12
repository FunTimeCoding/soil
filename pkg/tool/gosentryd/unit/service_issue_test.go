package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"testing"
)

func TestUpdateIssueChangesTheStoredStatus(t *testing.T) {
	c := mock_client.New()
	c.AddIssue("100", "41", constant.UnresolvedStatus)
	s := service_tester.Service(c)
	_, e := s.UpdateIssue("100", constant.ResolvedStatus, "")
	assert.FatalOnError(t, e)
	issue, f := s.Issue("100")
	assert.FatalOnError(t, f)
	assert.String(t, "resolved", issue.Status)
}

func TestDeletedIssueIsNotFound(t *testing.T) {
	c := mock_client.New()
	c.AddIssue("100", "41", constant.UnresolvedStatus)
	c.AddIssue("200", "41", constant.UnresolvedStatus)
	s := service_tester.Service(c)
	assert.FatalOnError(t, s.DeleteIssue("100"))
	_, e := s.Issue("100")
	assert.True(t, not_found.Is(e))
	remaining, f := s.SearchIssues("", "", 0, "")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
	assert.String(t, "200", remaining[0].Identifier)
}

func TestUpdatingAMissingIssueIsNotFound(t *testing.T) {
	_, e := service_tester.Service(mock_client.New()).UpdateIssue(
		"999",
		constant.ResolvedStatus,
		"",
	)
	assert.True(t, not_found.Is(e))
}
