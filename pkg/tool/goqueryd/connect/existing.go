package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func Existing(
	c *client.Client,
	collection string,
) map[string]string {
	wrapped := &client.ClientWithResponses{ClientInterface: c}
	result := map[string]string{}
	offset := 0

	for {
		limit := constant.ListPage
		currentOffset := offset
		r, e := wrapped.GetListWithResponse(
			context.Background(),
			&client.GetListParams{
				Collection: collection,
				Limit:      &limit,
				Offset:     &currentOffset,
			},
		)
		errors.PanicOnError(e)

		if r.JSON200 == nil {
			return result
		}

		for _, entry := range r.JSON200.Results {
			result[entry.Path] = entry.Hash
		}

		if len(r.JSON200.Results) < constant.ListPage {
			return result
		}

		offset += constant.ListPage
	}
}
