package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) MustMetadata(metric string) map[string][]v1.Metadata {
	result, e := c.Metadata(metric)
	errors.PanicOnError(e)

	return result
}
