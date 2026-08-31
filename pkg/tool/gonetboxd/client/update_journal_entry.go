package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateJournalEntry(
	identifier int32,
	kind string,
	comments string,
) *response.Response {
	body := client.UpdateJournalEntryJSONRequestBody{}

	if comments != "" {
		body.Comments = &comments
	}

	if kind != "" {
		body.Kind = &kind
	}

	result, e := c.client.UpdateJournalEntry(c.context, identifier, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
