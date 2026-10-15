package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_record"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_request"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) CreateConfigurationContext(
	name string,
	weight int32,
	payload map[string]any,
	a *configuration_assignment.Assignment,
) (*configuration_context.Context, error) {
	q := configuration_request.New()
	q.Name = name
	q.Weight = &weight
	q.Payload = &payload

	if e := c.assignmentFields(q, a); e != nil {
		return nil, e
	}

	result := configuration_record.New()

	if f := c.rawRequest(
		http.MethodPost,
		constant.ConfigurationContextPath,
		q,
		result,
	); f != nil {
		return nil, f
	}

	return configuration_context.New(result), nil
}
