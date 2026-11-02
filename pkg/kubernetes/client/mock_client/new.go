package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"k8s.io/api/networking/v1"
)

func New() *Client {
	result := &Client{
		pods:     map[string]bool{},
		policies: map[string]*v1.NetworkPolicy{},
	}
	result.failures = failure.New(result.Certificates)

	return result
}
