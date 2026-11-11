package document

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func indexedSession(r *client.SearchResult) string {
	if r.Metadata == nil {
		return ""
	}

	value := (*r.Metadata)["session_id"]

	if len(value) == 0 {
		return ""
	}

	return value[0]
}
