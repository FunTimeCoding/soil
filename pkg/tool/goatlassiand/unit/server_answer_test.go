package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func TestUnknownIssueIsAnsweredNotFoundUncaptured(t *testing.T) {
	m, e := client_tester.GetIssue(t, not_found.Format("Issue does not exist"))
	assert.True(t, not_found.Is(e))
	assert.Count(t, 0, m.Events())
}

func TestRejectedQueryIsAnsweredBadRequestUncaptured(t *testing.T) {
	j := service_tester.NewJira(t)
	j.Fail(j.Search, validation.New("Error in the JQL Query"))
	c, m := client_tester.NewReported(t, j)
	_, e := c.SearchIssues(context.Background(), "bravo =", nil)
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, m.Events())
}
