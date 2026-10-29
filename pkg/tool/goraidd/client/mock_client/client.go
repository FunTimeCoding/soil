package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"sync"
)

type Client struct {
	mutex   sync.Mutex
	Logs    client.LogsResponse
	Reports []client.ReportResponse
	Failure error
	Probe   error
	Closed  int
}
