package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func TestRequiringFindsPinnedConsumer(t *testing.T) {
	library := t.TempDir()
	user := t.TempDir()
	other := t.TempDir()
	system_tester.New(t).Write(
		library,
		"go.mod",
		"module other.test/lib\n\ngo 1.26.3\n",
	)
	system_tester.New(t).Write(
		user,
		"go.mod",
		"module other.test/user\n\ngo 1.26.3\n\nrequire other.test/lib v0.1.0\n",
	)
	system_tester.New(t).Write(
		other,
		"go.mod",
		"module other.test/other\n\ngo 1.26.3\n",
	)
	i := inventory.New()
	i.Add("lib", library)
	i.Add("user", user)
	i.Add("other", other)
	assert.Strings(t, []string{user}, i.Requiring(library))
	assert.Count(t, 0, i.Replacing(library))
}
