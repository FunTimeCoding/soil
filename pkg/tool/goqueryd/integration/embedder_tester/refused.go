package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
)

func (o *Tester) Refused(
	e face.Embedder,
	name string,
	body string,
) {
	o.t.Helper()
	chunks := chunk.Document(body, name, service_tester.New(o.t).Reranker)
	assert.Count(o.t, 1, chunks)
	_, f := embed.Single(o.t.Context(), e, chunks[0].Text)
	assert.True(o.t, validation.Is(f))
}
