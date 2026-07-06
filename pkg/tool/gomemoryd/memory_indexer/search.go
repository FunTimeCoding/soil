package memory_indexer

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (i *Indexer) Search(o *search_option.Option) ([]face.SearchResult, error) {
	params := &client.GetSearchParams{
		Query:      o.Query,
		Collection: &o.Collection,
		Limit:      &o.Limit,
		Full:       new(true),
		Mode:       new(client.GetSearchParamsMode("hybrid")),
	}

	if len(o.Exclude) > 0 {
		params.Exclude = &o.Exclude
	}

	r, e := i.client.GetSearch(
		context.Background(),
		params,
	)

	if e != nil {
		return nil, e
	}

	parsed, f := client.ParseGetSearchResponse(r)

	if f != nil {
		return nil, f
	}

	if parsed.JSON200 == nil {
		return nil, fmt.Errorf(
			"unexpected response: %s",
			parsed.Status(),
		)
	}

	var result []face.SearchResult

	for _, s := range parsed.JSON200.Results {
		body := ""

		if s.Body != nil {
			body = *s.Body
		}

		result = append(
			result,
			face.SearchResult{
				Path:  s.Path,
				Body:  body,
				Score: float64(s.Score),
			},
		)
	}

	return result, nil
}
