package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/query/filter"
	"testing"
)

func TestQueryFilterNew(t *testing.T) {
	assert.NotNil(t, filter.New())
}
