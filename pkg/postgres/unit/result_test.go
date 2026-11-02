package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"testing"
)

func TestAResultKeepsDuplicateColumnsInOrder(t *testing.T) {
	r := result.New([]string{"id", "id"}, [][]any{{1, 7}, {2, 9}})
	assert.Count(t, 2, r.Columns)
	assert.Any(t, 7, r.Rows[0][1])
}
