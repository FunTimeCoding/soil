package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *RestClient) SaveImpression(
	content string,
	source string,
) {
	r, e := c.http.PostImpressions(
		context.Background(),
		client.PostImpressionsJSONRequestBody{
			Content: content,
			Source:  &source,
		},
	)

	if e != nil {
		return
	}

	errors.PanicClose(r.Body)
}
