package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewContext(cluster string) (*Client, error) {
	c, e := client_configuration.NewContext(cluster)

	if e != nil {
		return nil, e
	}

	result, f := FromConfiguration(c, cluster)

	if f != nil {
		return nil, f
	}

	_, g := operation.Node(result.client).List(
		result.context,
		v1.ListOptions{},
	)

	if g != nil {
		return nil, g
	}

	return result, nil
}
