package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/rechunk_tester"
	"testing"
)

func TestRechunkLeavesMatchingEmbeddings(t *testing.T) {
	v, s := rechunk_tester.New(t, 0)
	documents, e := v.Rechunk()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, documents)
	assert.Count(t, 0, s.PendingEmbeddings())
}

func TestRechunkDeletesEmbeddingsAtStalePositions(t *testing.T) {
	v, s := rechunk_tester.New(t, 7)
	documents, e := v.Rechunk()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, documents)
	assert.String(t, "test/alfa.md", documents[0])
	assert.Count(t, 1, s.PendingEmbeddings())
}
