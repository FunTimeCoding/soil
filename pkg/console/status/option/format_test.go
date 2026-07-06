package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	alfa := New()
	alfa.Tags = []string{"a"}
	bravo := alfa.Copy()
	bravo.Tags = append(bravo.Tags, "b")
	assert.Strings(t, []string{"a"}, alfa.Tags)
	assert.Strings(t, []string{"a", "b"}, bravo.Tags)

	if alfa == bravo {
		t.Errorf(
			"Expect copy to be different: %+v %+v",
			alfa,
			bravo,
		)
	}
}
