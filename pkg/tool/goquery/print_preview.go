package goquery

import (
	"context"
	"fmt"
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
	printPreviewSections(t, p.Sections, p.Allowance)
	t.Line("windows")
	over := 0

	for _, k := range p.Chunks {
		over += printPreviewChunk(t, k, p.Allowance, p.EmbedAllowance)
	}

	t.Line()
	t.Line(
		fmt.Sprintf(
			"%d chunks, %d over %d tokens (%s, window %d) or the embedder's %d",
			len(p.Chunks),
			over,
			p.Allowance,
			p.Model,
			p.Window,
			p.EmbedAllowance,
		),
	)

	return over
}
