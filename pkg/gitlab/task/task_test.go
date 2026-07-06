package task

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestTask(t *testing.T) {
	a := New(&gitlab.Todo{})
	a.Raw = nil
	assert.Any(t, &Task{}, a)
}
