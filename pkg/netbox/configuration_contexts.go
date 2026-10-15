package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_record"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) ConfigurationContexts() ([]*configuration_context.Context, error) {
	var result configuration_record.Page

	if e := c.rawRequest(
		http.MethodGet,
		constant.ConfigurationContextListPath,
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return configuration_context.NewSlice(result.Results), nil
}
