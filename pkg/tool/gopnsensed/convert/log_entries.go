package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func LogEntries(v []*log_entry.Entry) []server.LogEntry {
	result := []server.LogEntry{}

	for _, e := range v {
		result = append(result, *LogEntry(e))
	}

	return result
}
