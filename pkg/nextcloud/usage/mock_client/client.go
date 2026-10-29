package mock_client

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"sync"
)

type Client struct {
	mutex   sync.Mutex
	usage   *usage.Usage
	Failure error
	Closed  int
}
