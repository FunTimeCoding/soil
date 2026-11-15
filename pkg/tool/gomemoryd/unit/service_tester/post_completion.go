package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (o *Tester) PostCompletion(
	path string,
	completedAt string,
) {
	o.t.Helper()
	metadata := map[string][]string{"source_type": {"session-completion"}}

	if completedAt != "" {
		metadata["completed_at"] = []string{completedAt}
	}

	_, e := o.Indexer.PostDocument(
		o.t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "completions",
			Path:       path,
			Body:       path,
			Metadata:   &metadata,
		},
	)
	assert.FatalOnError(o.t, e)
}
