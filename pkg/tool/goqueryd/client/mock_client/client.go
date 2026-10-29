package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"sync"
)

type Client struct {
	mutex     sync.Mutex
	documents map[string]*client.SearchResult
	Pushed    []client.PostDocumentJSONRequestBody
	Deleted   []client.DeleteDocumentParams
	Failure   error
	Closed    int
}
