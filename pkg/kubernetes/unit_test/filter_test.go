package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.NotNil(t, filter.New())
}
