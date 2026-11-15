package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"testing"
)

func New(t *testing.T) face.ReachableEmbedder {
	t.Helper()
	result := embed.NewEnvironment()
	t.Cleanup(result.Close)
	assert.SkipOnOutage(t, result.Reachable(t.Context()))

	return result
}
