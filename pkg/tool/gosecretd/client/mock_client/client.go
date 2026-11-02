package mock_client

import "sync"

type Client struct {
	mutex   sync.Mutex
	Secrets map[string]map[string]string
	Tamper  string
	Failure error
	Closed  int
}
