package gosentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func printEventEntries(entries []response.EventEntry) {
	for _, entry := range entries {
		switch entry.Type {
		case "exception":
			printException(entry.Payload.Values)
		case "message":
			m := entry.Payload.Formatted

			if m == "" {
				m = entry.Payload.Message
			}

			if m != "" {
				fmt.Printf("Message: %s\n", m)
			}
		}
	}
}
