package indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
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

	r, e := i.client.PostDocument(
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

	return r.Body.Close()
}
