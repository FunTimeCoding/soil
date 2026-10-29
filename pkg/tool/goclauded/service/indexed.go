package service

import (
	"context"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) indexed(collection string) (map[string]string, error) {
	result := map[string]string{}
	offset := 0

	for {
		limit := queryConstant.ListPage
		page := offset
		outcome, e := s.query.GetList(
			context.Background(),
			&client.GetListParams{
				Collection: collection,
				Limit:      &limit,
				Offset:     &page,
			},
		)

		if e != nil {
			return nil, e
		}

		for _, r := range outcome.Results {
			result[r.Path] = r.Hash
		}

		if len(outcome.Results) < queryConstant.ListPage {
			return result, nil
		}

		offset += queryConstant.ListPage
	}
}
