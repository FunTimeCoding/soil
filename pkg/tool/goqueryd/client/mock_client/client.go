package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"sync"
)

type Client struct {
	mutex              sync.Mutex
	documents          map[string]*client.SearchResult
	Pushed             []client.PostDocumentJSONRequestBody
	Deleted            []client.DeleteDocumentParams
	Indexed            []client.PostIndexJSONRequestBody
	Embedded           int
	Rechunked          int
	Tagged             []client.PostTagJSONRequestBody
	CollectionsAdded   []client.PostCollectionJSONRequestBody
	CollectionsDeleted []client.DeleteCollectionParams
	ContextsAdded      []client.PostContextJSONRequestBody
	ContextsDeleted    []client.DeleteContextParams
	Failure            error
	Closed             int
}
