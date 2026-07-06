package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func formatCheckContext(r *client.CheckResponse) string {
	if len(r.Entries) == 0 {
		return ""
	}

	groups := map[string][]string{}

	for _, e := range r.Entries {
		groups[e.Kind] = append(groups[e.Kind], e.Body)
	}

	var parts []string

	if v, okay := groups[constant.QueueReannounce]; okay {
		parts = append(parts, v...)
	}

	if v, okay := groups[constant.QueuePulse]; okay {
		for _, body := range v {
			parts = append(parts, fmt.Sprintf("[pulse] %s", body))
		}
	}

	var sessionActivity []string

	for _, kind := range []string{
		constant.QueueSessionAnnounce,
		constant.QueueSessionRelease,
		constant.QueueSessionComplete,
		constant.QueueSessionUpdate,
	} {
		if v, okay := groups[kind]; okay {
			sessionActivity = append(sessionActivity, v...)
		}
	}

	if len(sessionActivity) > 0 {
		parts = append(parts, "Session activity:")

		for _, body := range sessionActivity {
			parts = append(parts, fmt.Sprintf("  %s", body))
		}
	}

	var memoryActivity []string

	for _, kind := range []string{
		constant.QueueMemoryCreate,
		constant.QueueMemoryUpdate,
	} {
		if v, okay := groups[kind]; okay {
			memoryActivity = append(memoryActivity, v...)
		}
	}

	if len(memoryActivity) > 0 {
		parts = append(parts, "Memory activity:")

		for _, body := range memoryActivity {
			parts = append(parts, fmt.Sprintf("  %s", body))
		}
	}

	if v, okay := groups[constant.QueueTimeout]; okay {
		for _, body := range v {
			parts = append(parts, fmt.Sprintf("Idle: %s", body))
		}
	}

	if v, okay := groups[constant.QueueMessage]; okay {
		parts = append(parts, "Messages:")

		for _, body := range v {
			parts = append(parts, fmt.Sprintf("  %s", body))
		}
	}

	if v, okay := groups[constant.QueueNotification]; okay {
		parts = append(parts, "Notifications:")

		for _, body := range v {
			parts = append(parts, fmt.Sprintf("  %s", body))
		}
	}

	return join.NewLine(parts)
}
