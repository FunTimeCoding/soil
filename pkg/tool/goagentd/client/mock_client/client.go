package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"sync"
)

type Client struct {
	mutex   sync.Mutex
	Status  client.StatusResponse
	Intents []string
	Failure error
	Probe   error
	Asked   int
	Closed  int
}
