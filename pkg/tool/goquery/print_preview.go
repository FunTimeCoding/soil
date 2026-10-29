package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func printPreview(
	c face.Client,
	t *terminal.Terminal,
	name string,
	body string,
) int {
	p, e := c.PostChunk(
		context.Background(),
		client.PostChunkJSONRequestBody{Path: name, Body: body},
	)
	t.EmitDone(e)
	printPreviewSections(p.Sections, p.Allowance)
	console.Format("windows\n")
	over := 0

	for _, k := range p.Chunks {
		over += printPreviewChunk(k, p.Allowance)
	}

	console.Format(
		"\n%d chunks, %d over %d tokens (%s, window %d)\n",
		len(p.Chunks),
		over,
		p.Allowance,
		p.Model,
		p.Window,
	)

	return over
}
