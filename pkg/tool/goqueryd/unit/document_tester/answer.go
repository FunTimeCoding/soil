package document_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
)

func (o *Tester) Answer(
	path string,
) (*client.GetDocumentResponse, *memory.Memory) {
	o.t.Helper()
	q := service_tester.New(o.t)
	c := client_tester.GeneratedClient(o.t, q.Served())
	o.Push(c, "notes", "alfa.md", "# Alfa\n\nzulu.\n")
	answer, e := c.GetDocumentWithResponse(
		o.t.Context(),
		&client.GetDocumentParams{Path: path},
	)
	assert.FatalOnError(o.t, e)

	return answer, q.Reporter
}
