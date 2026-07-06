package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func Tasks(v []*task.Task) []*server.Task {
	result := make([]*server.Task, 0, len(v))

	for _, t := range v {
		result = append(result, Task(t))
	}

	return result
}
