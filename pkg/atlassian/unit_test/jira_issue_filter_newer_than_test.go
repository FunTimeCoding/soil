package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
	"time"
)

func TestFilterNewerThan(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Created = jira.Time(time.Now().Add(-3 * time.Hour))
	r2 := issue.Raw("TEST-2")
	r2.Fields.Created = jira.Time(time.Now())
	actual := issue.FilterNewerThan(
		[]*issue.Issue{issue.New(r1, o), issue.New(
			r2,
			o,
		)},
		2,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
