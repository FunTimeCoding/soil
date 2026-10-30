package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func printFailed(r *client.FailedReport) {
	for _, c := range r.Chunks {
		console.Format(
			"%s/%s  chunk %d  lines %d-%d  %d B\n  %s\n",
			c.Collection,
			c.Path,
			c.Index,
			c.FirstLine,
			c.LastLine,
			c.Bytes,
			c.Reason,
		)
	}

	console.Format("\n%d chunks refused by the embedder\n", len(r.Chunks))
}
