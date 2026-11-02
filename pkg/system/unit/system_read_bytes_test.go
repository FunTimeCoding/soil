package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
)

func TestReadBytesLeavesNoHandleOpen(t *testing.T) {
	base := t.TempDir()
	assert.FatalOnError(
		t,
		os.WriteFile(filepath.Join(base, "alfa"), []byte("bravo"), 0o600),
	)
	before := fixture.OpenHandles(t)

	for range 200 {
		assert.String(t, "bravo", string(system.ReadBytes(base, "alfa")))
	}

	assert.True(t, fixture.OpenHandles(t)-before < 50)
}
