package issue

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFilterLabel(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Labels = []string{upper.Alfa}
	r2 := Raw("TEST-2")
	r2.Fields.Labels = []string{upper.Bravo}
	actual := FilterLabel([]*Issue{New(r1, o), New(r2, o)}, upper.Bravo)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
