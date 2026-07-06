package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"testing"
)

func TestGroupByStatus(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Status = &jira.Status{Name: constant.ToDo}
	r2 := Raw("TEST-2")
	r2.Fields.Status = &jira.Status{Name: constant.Closed}
	actual := GroupByStatus([]*Issue{New(r1, o), New(r2, o)})
	assert.Count(t, 2, actual)
	assert.Count(t, 1, actual[constant.ToDo])
	assert.String(t, "TEST-1", actual[constant.ToDo][0].Key)
	assert.Count(t, 1, actual[constant.Closed])
	assert.String(t, "TEST-2", actual[constant.Closed][0].Key)
}
