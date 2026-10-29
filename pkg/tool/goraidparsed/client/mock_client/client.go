package mock_client

import "sync"

type Client struct {
	mutex     sync.Mutex
	Generated [][]string
	Calls     int
	Failure   error
	Probe     error
	Closed    int
}
