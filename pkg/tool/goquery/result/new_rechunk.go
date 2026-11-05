package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func NewRechunk(
	queued []string,
	embedded *client.EmbedResult,
) *Rechunk {
	return &Rechunk{Queued: queued, Embedded: embedded}
}
