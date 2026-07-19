package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestTask(t *testing.T) {
	a := task.New(&gitlab.Todo{})
	a.Raw = nil
	assert.Any(t, &task.Task{}, a)
}
