package indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (i *Indexer) Delete(
	collection string,
	path string,
) error {
	r, e := i.client.DeleteDocumentWithResponse(
		context.Background(),
		&client.DeleteDocumentParams{Collection: collection, Path: path},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("delete document", r.HTTPResponse, r.Body)
	}

	return nil
}
