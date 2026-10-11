package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"sync"
	"testing"
)

func TestWorkspaceReferencesAreSafeInParallel(t *testing.T) {
	w := index.New(t.TempDir(), kind_tester.Module(t))
	var group sync.WaitGroup
	results := make([]any, 8)

	for i := range results {
		group.Go(
			func() {
				results[i] = w.References()
			},
		)
	}

	group.Wait()

	for _, r := range results {
		assert.True(t, r == results[0])
	}
}
