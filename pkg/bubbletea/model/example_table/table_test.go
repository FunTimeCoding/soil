package example_table

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestTable(t *testing.T) {
	assert.NotNil(t, New(&table.Model{}))
}
