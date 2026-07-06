package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestFilterNewerThan(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Created = jira.Time(time.Now().Add(-3 * time.Hour))
	r2 := Raw("TEST-2")
	r2.Fields.Created = jira.Time(time.Now())
	actual := FilterNewerThan([]*Issue{New(r1, o), New(r2, o)}, 2)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
