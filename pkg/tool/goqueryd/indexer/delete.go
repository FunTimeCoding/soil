package indexer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (i *Indexer) Delete(collection string, path string) error {
	r, e := i.client.DeleteDocument(
		context.Background(),
		&client.DeleteDocumentParams{Collection: collection, Path: path},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}
