package mock_client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
)

func (c *Client) MustNamespaces(_ *filter.Filter) []*namespace.Namespace {
	return nil
}
