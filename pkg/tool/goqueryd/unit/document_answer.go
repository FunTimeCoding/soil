package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func documentAnswer(
	t *testing.T,
	path string,
) (*client.GetDocumentResponse, *memory.Memory) {
	t.Helper()
	r := memory.New()
	v := service_tester.New(
		t,
		service_tester.NewOption(t, newEmbedder(t), newReranker(t)),
		r,
	)
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	answer, e := mountedClient(t, v, r).GetDocumentWithResponse(
		t.Context(),
		&client.GetDocumentParams{Path: path},
	)
	assert.FatalOnError(t, e)

	return answer, r
}
