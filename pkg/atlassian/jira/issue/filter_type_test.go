package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFilterType(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Type = jira.IssueType{Name: EpicType}
	r2 := Raw("TEST-2")
	r2.Fields.Type = jira.IssueType{Name: TaskType}
	actual := FilterType([]*Issue{New(r1, o), New(r2, o)}, TaskType)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
