package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func Existing(
	c *client.ClientWithResponses,
	collection string,
) (map[string]string, error) {
	result := map[string]string{}
	offset := 0

	for {
		limit := constant.ListPage
		currentOffset := offset
		r, e := c.GetListWithResponse(
			context.Background(),
			&client.GetListParams{
				Collection: collection,
				Limit:      &limit,
				Offset:     &currentOffset,
			},
		)

		if e != nil {
			return nil, e
		}

		if r.JSON200 == nil {
			return nil, refusal.New("list documents", r.HTTPResponse, r.Body)
		}

		for _, entry := range r.JSON200.Results {
			result[entry.Path] = entry.Hash
		}

		if len(r.JSON200.Results) < constant.ListPage {
			return result, nil
		}

		offset += constant.ListPage
	}
}
