package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_record"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) updateConfigurationContext(
	v *configuration_context.Context,
	q *configuration_request.Request,
) (*configuration_context.Context, error) {
	result := configuration_record.New()

	if e := c.rawRequest(
		http.MethodPatch,
		fmt.Sprintf("%s%d/", constant.ConfigurationContextPath, v.Identifier),
		q,
		result,
	); e != nil {
		return nil, e
	}

	return configuration_context.New(result), nil
}
