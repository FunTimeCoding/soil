package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

type Client struct {
	statuses  map[string]*client.Status
	failures  *failure.Failure
	Requested []client.PowerRequest
	Failure   error
	Closed    int
}
