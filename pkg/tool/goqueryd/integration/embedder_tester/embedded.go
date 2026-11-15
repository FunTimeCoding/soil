package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Embedded(
	t *testing.T,
	e face.Embedder,
	name string,
	body string,
) {
	t.Helper()

	for _, c := range chunk.Document(body, name, service_tester.NewReranker(t)) {
		_, f := embed.Single(t.Context(), e, c.Text)
		assert.FatalOnError(t, f)
	}
}
