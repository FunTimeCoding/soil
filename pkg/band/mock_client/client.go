package mock_client

import "sync"

type Client struct {
	state         int
	available     []int
	requested     []int
	listener      bool
	screenService int
	writes        []string
	failure       error
	Closed        int
	mutex         sync.Mutex
}
