package mock_client

import (
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/postgres/types/statement"
	"sync"
)

type Client struct {
	mutex      sync.Mutex
	results    map[string]*result.Result
	statements []*statement.Statement
	Failure    error
	Closed     int
}
