package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"testing"
)

func TestWithProjects(t *testing.T) {
	assert.NotNil(t, gitlab.WithProjects([]int64{}))
}
