package mock_client

import "sync"

type Client struct {
	mutex   sync.Mutex
	Failure error
	Closed  int
}
