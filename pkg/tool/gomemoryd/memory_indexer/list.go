package memory_indexer

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (i *Indexer) List(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) (*face.ListOutcome, error) {
	params := &client.GetListParams{
		Collection: collection,
		Full:       &full,
	}

	if len(metadata) > 0 {
		params.Metadata = &metadata
	}

	if limit > 0 {
		params.Limit = &limit
	}

	if offset > 0 {
		params.Offset = &offset
	}

	r, e := i.client.GetList(context.Background(), params)

	if e != nil {
		return nil, e
	}

	parsed, e := client.ParseGetListResponse(r)

	if e != nil {
		return nil, e
	}

	if parsed.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response: %s", parsed.Status())
	}

	var results []face.SearchResult

	for _, s := range parsed.JSON200.Results {
		body := ""

		if s.Body != nil {
			body = *s.Body
		}

		results = append(
			results,
			face.SearchResult{
				Path:  s.Path,
				Body:  body,
				Score: float64(s.Score),
			},
		)
	}

	return face.NewListOutcome(results), nil
}
