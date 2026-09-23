package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/status"
	"time"
)

func Process(v *status.Status) *server.Process {
	result := &server.Process{
		Name:    v.Name,
		Command: v.Command,
		Running: v.Running,
	}

	if v.StartedAt == "" {
		return result
	}

	if started, e := time.Parse(time.RFC3339, v.StartedAt); e == nil {
		result.Started = &started
	}

	return result
}
