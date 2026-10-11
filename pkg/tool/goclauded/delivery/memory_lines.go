package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func memoryLines(
	groups map[string][]queue.Entry,
	room int,
) []string {
	var entries []queue.Entry
	entries = append(entries, groups[constant.QueueMemoryCreate]...)
	entries = append(entries, groups[constant.QueueMemoryUpdate]...)
	lines := bodies(entries, constant.DeliveryIndent)

	if len(lines) == 0 {
		return nil
	}

	result := section(constant.DeliveryMemoryActivity, lines)

	if size(result) <= room {
		return result
	}

	used := lineSize(constant.DeliveryMemoryActivity) + lineSize(trimLine(len(lines)))
	kept := 0

	for kept < len(lines) && used+lineSize(lines[kept]) <= room {
		used += lineSize(lines[kept])
		kept++
	}

	result = []string{constant.DeliveryMemoryActivity}
	result = append(result, lines[:kept]...)

	return append(result, trimLine(len(lines)-kept))
}
