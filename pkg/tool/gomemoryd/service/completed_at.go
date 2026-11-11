package service

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"time"
)

func completedAt(r *client.SearchResult) time.Time {
	if r.Metadata == nil {
		return time.Time{}
	}

	value := (*r.Metadata)["completed_at"]

	if len(value) == 0 {
		return time.Time{}
	}

	return soilTime.Parse(time.RFC3339, value[0])
}
