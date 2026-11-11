package document

import (
	"context"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) indexed(
	collection string,
) (map[string]client.SearchResult, error) {
	result := map[string]client.SearchResult{}
	offset := 0

	for {
		limit := goquerydConstant.ListPage
		page := offset
		var outcome *client.ListOutcome

		if e := s.queryGuard.Attempt(
			context.Background(),
			func(x context.Context) error {
				var f error
				outcome, f = s.query.GetList(
					x,
					&client.GetListParams{
						Collection: collection,
						Limit:      &limit,
						Offset:     &page,
					},
				)

				return f
			},
		); e != nil {
			return nil, e
		}

		for _, r := range outcome.Results {
			result[r.Path] = r
		}

		if len(outcome.Results) < goquerydConstant.ListPage {
			return result, nil
		}

		offset += goquerydConstant.ListPage
	}
}
