package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"slices"
)

func (c *Client) MergeRequests(
	_ context.Context,
	f *merge_request_filter.Filter,
) ([]*merge_request.Request, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	var result []*merge_request.Request

	for _, r := range c.requests {
		if matches(r, f) {
			result = append(result, r)
		}
	}

	slices.SortStableFunc(
		result,
		func(
			a *merge_request.Request,
			b *merge_request.Request,
		) int {
			if f.OrderBy == constant.CreatedOrder {
				return moment(b.Create).Compare(moment(a.Create))
			}

			return moment(b.Update).Compare(moment(a.Update))
		},
	)

	if f.Offset >= len(result) {
		return nil, nil
	}

	return result[f.Offset:min(f.Offset+f.Limit, len(result))], nil
}
