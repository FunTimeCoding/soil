package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/indexer"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newIndexer(
	t *testing.T,
	status int,
	answer string,
) *indexer.Indexer {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	c, e := client.NewClientWithResponses(
		s.URL,
		client.WithHTTPClient(s.Client()),
	)
	assert.FatalOnError(t, e)

	return indexer.New(c, "alfa")
}
