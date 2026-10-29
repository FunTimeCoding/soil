package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func TestAnEmbeddingRecordsTheStoreModel(t *testing.T) {
	d := connection.NewMemory()
	defer errors.PanicClose(d)
	s := store.New(d, "nomic-embed-text-v1.5.f16")
	s.InsertEmbedding("alfa", 0, 0, []float32{1, 0}, "2024-01-01T00:00:00Z")
	var model string
	assert.FatalOnError(
		t,
		d.QueryRow(`SELECT model FROM embedding WHERE hash = 'alfa'`).Scan(
			&model,
		),
	)
	assert.String(t, "nomic-embed-text-v1.5.f16", model)
}
