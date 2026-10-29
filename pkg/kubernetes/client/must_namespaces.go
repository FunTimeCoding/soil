package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
)

func (c *Client) MustNamespaces(f *filter.Filter) []*namespace.Namespace {
	result, e := c.Namespaces(f)
	errors.PanicOnError(e)

	return result
}
