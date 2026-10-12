package mock_client

import "sync"

type Client struct {
	state     int
	available []int
	requested []int
	mutex     sync.Mutex
}
