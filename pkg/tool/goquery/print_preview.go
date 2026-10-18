package goquery

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func printPreview(
	c *client.ClientWithResponses,
	name string,
	body string,
) int {
	p, e := c.PostChunkWithResponse(
		context.Background(),
		client.PostChunkJSONRequestBody{Path: name, Body: body},
	)
	errors.PanicOnError(e)
	refusal.PanicOnFailure("chunk", p.HTTPResponse, p.Body)
	printPreviewSections(p.JSON200.Sections, p.JSON200.Allowance)
	console.Format("windows\n")
	over := 0

	for _, k := range p.JSON200.Chunks {
		over += printPreviewChunk(k, p.JSON200.Allowance)
	}

	console.Format(
		"\n%d chunks, %d over %d tokens (%s, window %d)\n",
		len(p.JSON200.Chunks),
		over,
		p.JSON200.Allowance,
		p.JSON200.Model,
		p.JSON200.Window,
	)

	return over
}
