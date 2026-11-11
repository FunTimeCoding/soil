package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func postCompletion(
	t *testing.T,
	o *service_tester.Tester,
	path string,
	completedAt string,
) {
	t.Helper()
	metadata := map[string][]string{"source_type": {"session-completion"}}

	if completedAt != "" {
		metadata["completed_at"] = []string{completedAt}
	}

	_, e := o.Indexer.PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "completions",
			Path:       path,
			Body:       path,
			Metadata:   &metadata,
		},
	)
	assert.FatalOnError(t, e)
}
