package memory_indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
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

	r, e := i.client.GetSearchWithResponse(context.Background(), params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("search", r.HTTPResponse, r.Body)
	}

	var result []face.SearchResult

	for _, s := range r.JSON200.Results {
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
