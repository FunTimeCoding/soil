package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"sync"
)

type Client struct {
	mutex     sync.Mutex
	Base      string
	UserList  []*client.User
	GroupList []*client.Group
	Passwords map[string]string
	Failure   error
	Closed    int
	failures  *failure.Failure
}
