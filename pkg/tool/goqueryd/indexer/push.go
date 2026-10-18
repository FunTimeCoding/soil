package indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (i *Indexer) Push(
	collection string,
	path string,
	body string,
	metadata map[string][]string,
) error {
	merged := map[string][]string{constant.SourceType: {i.sourceType}}

	for key, value := range metadata {
		merged[key] = value
	}

	r, e := i.client.PostDocumentWithResponse(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: collection,
			Path:       path,
			Body:       body,
			Metadata:   &merged,
		},
	)

	if e != nil {
		return e
	}

	if r.JSON200 == nil {
		return refusal.New("push document", r.HTTPResponse, r.Body)
	}

	return nil
}
