package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"testing"
)

func TestFilterNarrows(t *testing.T) {
	r := registry.NewDefault().Filter([]string{"go", "yaml"})
	assert.Count(t, 2, r.Languages())
	assert.String(t, "Go", r.ByPath("a.go").Name)
	assert.String(t, "Go", r.ByPath("a_test.go").Name)
	assert.Nil(t, r.ByPath("a.php"))
}

func TestFilterEmptyKeepsAll(t *testing.T) {
	r := registry.NewDefault()
	assert.Count(t, len(r.Languages()), r.Filter(nil).Languages())
}

func TestFilterUnknownPanics(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	registry.NewDefault().Filter([]string{"cobol"})
}
