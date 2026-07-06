package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestSortByAge(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Created = jira.Time(time.Now().Add(-3 * time.Hour))
	r2 := Raw("TEST-2")
	r2.Fields.Created = jira.Time(time.Now().Add(-2 * time.Hour))
	r3 := Raw("TEST-3")
	r3.Fields.Created = jira.Time(time.Now().Add(-1 * time.Hour))
	actual := SortByAge([]*Issue{New(r3, o), New(r1, o), New(r2, o)})
	assert.Count(t, 3, actual)
	assert.String(t, "TEST-1", actual[0].Key)
	assert.String(t, "TEST-2", actual[1].Key)
	assert.String(t, "TEST-3", actual[2].Key)
}
