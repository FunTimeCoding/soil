package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
	"time"
)

func TestSortByAge(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Created = jira.Time(time.Now().Add(-3 * time.Hour))
	r2 := issue.Raw("TEST-2")
	r2.Fields.Created = jira.Time(time.Now().Add(-2 * time.Hour))
	r3 := issue.Raw("TEST-3")
	r3.Fields.Created = jira.Time(time.Now().Add(-1 * time.Hour))
	actual := issue.SortByAge([]*issue.Issue{issue.New(r3, o), issue.New(r1, o), issue.New(
		r2,
		o,
	)})
	assert.Count(t, 3, actual)
	assert.String(t, "TEST-1", actual[0].Key)
	assert.String(t, "TEST-2", actual[1].Key)
	assert.String(t, "TEST-3", actual[2].Key)
}
