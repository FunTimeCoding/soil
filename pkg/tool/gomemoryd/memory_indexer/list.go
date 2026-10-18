package memory_indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (i *Indexer) List(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) (*face.ListOutcome, error) {
	params := &client.GetListParams{Collection: collection, Full: &full}

	if len(metadata) > 0 {
		params.Metadata = &metadata
	}

	if limit > 0 {
		params.Limit = &limit
	}

	if offset > 0 {
		params.Offset = &offset
	}

	r, e := i.client.GetListWithResponse(context.Background(), params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list documents", r.HTTPResponse, r.Body)
	}

	var results []face.SearchResult

	for _, s := range r.JSON200.Results {
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
