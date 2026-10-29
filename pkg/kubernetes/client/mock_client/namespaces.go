package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
)

func (c *Client) Namespaces(_ *filter.Filter) ([]*namespace.Namespace, error) {
	return nil, c.Failure
}
