package gitlab

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestWithProjects(t *testing.T) {
	assert.NotNil(t, WithProjects([]int64{}))
}
