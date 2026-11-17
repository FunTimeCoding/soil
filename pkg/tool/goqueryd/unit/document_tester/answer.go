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
	q := service_tester.New(t)
	c := client_tester.GeneratedClient(t, base.Serve(t, q))
	Push(t, c, "notes", "alfa.md", "# Alfa\n\nzulu.\n")
	answer, e := c.GetDocumentWithResponse(
		t.Context(),
		&client.GetDocumentParams{Path: path},
	)
	assert.FatalOnError(t, e)

	return answer, q.Reporter
}
