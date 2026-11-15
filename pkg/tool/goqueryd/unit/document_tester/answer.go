package document_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Answer(
	t *testing.T,
	path string,
) (*client.GetDocumentResponse, *memory.Memory) {
	t.Helper()
	r := memory.New()
	c := client_tester.Rest(
		t,
		base.Serve(
			t,
			service_tester.NewOption(
				t,
				service_tester.NewEmbedder(t),
				service_tester.NewReranker(t),
			),
			r,
		),
	)
	Push(t, c, "notes", "alfa.md", "# Alfa\n\nzulu.\n")
	answer, e := c.GetDocumentWithResponse(
		t.Context(),
		&client.GetDocumentParams{Path: path},
	)
	assert.FatalOnError(t, e)

	return answer, r
}
