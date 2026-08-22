package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
)

func (c *Client) DeleteSilence(identifier string) error {
	p := silence.NewDeleteSilenceParams()
	p.SilenceID = strfmt.UUID(identifier)
	result, e := c.client.Silence.DeleteSilence(p)

	if e != nil {
		return e
	}

	if !result.IsSuccess() {
		return unexpected.Format("unexpected response: %+v", result)
	}

	return nil
}
