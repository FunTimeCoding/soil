package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"sync"
)

type Client struct {
	mutex    sync.Mutex
	ViewList []client.View
	Failure  error
	Closed   int
	next     int
	failures *failure.Failure
}
