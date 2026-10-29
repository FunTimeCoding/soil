package mock_client

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
)

func (c *Client) Dynamic() dynamic.Interface {
	return fake.NewSimpleDynamicClient(runtime.NewScheme())
}
