package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"testing"
)

func TestFilterStatus(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Status = &jira.Status{Name: constant.ToDo}
	r2 := Raw("TEST-2")
	r2.Fields.Status = &jira.Status{Name: constant.Closed}
	actual := FilterStatus([]*Issue{New(r1, o), New(r2, o)}, constant.Closed)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
