package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
)

func Statuses(v []*status.Status) []*server.Status {
	result := make([]*server.Status, len(v))

	for i, s := range v {
		result[i] = Status(s)
	}

	return result
}
