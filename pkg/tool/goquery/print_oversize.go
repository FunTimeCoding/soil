package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func printOversize(r *client.OversizeReport) {
	chunks := 0

	for _, f := range r.Files {
		console.Format(
			"%s/%s  %d chunks, worst %d/%d\n",
			f.Collection,
			f.Path,
			len(f.Chunks),
			f.Worst,
			r.Allowance,
		)

		for _, c := range f.Chunks {
			console.Format(
				"  chunk %d  lines %d-%d  %d B  %d tokens  over by %d\n",
				c.Index,
				c.FirstLine,
				c.LastLine,
				c.Bytes,
				c.Tokens,
				c.Tokens-r.Allowance,
			)
		}

		chunks += len(f.Chunks)
	}

	console.Format(
		"\n%d files, %d chunks over %d tokens (%s, window %d)\n",
		len(r.Files),
		chunks,
		r.Allowance,
		r.Model,
		r.Window,
	)
}
