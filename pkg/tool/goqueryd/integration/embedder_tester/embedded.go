package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
)

func (o *Tester) Embedded(
	e face.Embedder,
	name string,
	body string,
) {
	o.t.Helper()

	for _, c := range chunk.Document(
		body,
		name,
		service_tester.New(o.t).Reranker,
	) {
		_, f := embed.Single(o.t.Context(), e, c.Text)
		assert.FatalOnError(o.t, f)
	}
}
