package unit_test

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_table"
	"testing"
)

func TestTable(t *testing.T) {
	assert.NotNil(t, example_table.New(&table.Model{}))
}
