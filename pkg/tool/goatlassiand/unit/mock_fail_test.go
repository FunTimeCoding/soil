package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"testing"
)

func TestMockIssueFailsOnItsOwn(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	m.Found = []*client.JiraIssue{{Key: "ALFA-1"}}
	m.Fail(m.GetIssue, unavailable.Format("goatlassiand: down"))
	_, e := m.GetIssue(context.Background(), "ALFA-1", true)
	assert.True(t, unavailable.Is(e))
	found, f := m.SearchIssues(context.Background(), "", nil)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, found)
}
