package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func bodies(
	entries []queue.Entry,
	format string,
) []string {
	var result []string

	for _, e := range entries {
		result = append(result, fmt.Sprintf(format, e.Body))
	}

	return result
}
