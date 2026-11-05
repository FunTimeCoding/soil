package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

type Rechunk struct {
	Queued   []string            `json:"queued"`
	Embedded *client.EmbedResult `json:"embedded"`
}
