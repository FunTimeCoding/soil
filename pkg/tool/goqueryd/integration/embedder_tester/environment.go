package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
)

func (o *Tester) Environment() face.ReachableEmbedder {
	o.t.Helper()
	result := embed.NewEnvironment()
	o.t.Cleanup(result.Close)
	assert.SkipOnOutage(o.t, result.Reachable(o.t.Context()))

	return result
}
