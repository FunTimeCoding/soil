package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestSortByScore(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Created = jira.Time(time.Now().Add(-3 * time.Hour))
	i1 := New(r1, o)
	i1.SetScore(10)
	r2 := Raw("TEST-2")
	r2.Fields.Created = jira.Time(time.Now().Add(-2 * time.Hour))
	i2 := New(r2, o)
	i2.SetScore(5)
	r3 := Raw("TEST-3")
	r3.Fields.Created = jira.Time(time.Now().Add(-1 * time.Hour))
	i3 := New(r3, o)
	i3.SetScore(1)
	actual := SortByScore([]*Issue{i3, i1, i2})
	assert.Count(t, 3, actual)
	assert.String(t, "TEST-1", actual[0].Key)
	assert.String(t, "TEST-2", actual[1].Key)
	assert.String(t, "TEST-3", actual[2].Key)
}
