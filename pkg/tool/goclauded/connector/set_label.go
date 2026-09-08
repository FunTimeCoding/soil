package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
)

func (c *Client) SetLabel(
	sessionIdentifier string,
	key string,
	value string,
	from string,
) error {
	response, e := c.generated.PostSessionLabelWithResponse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionLabelJSONRequestBody{
			Key:   key,
			Value: &value,
			From:  &from,
		},
	)

	if e != nil {
		return e
	}

	if response.StatusCode() != http.StatusOK {
		return unexpected.Format("label status %d", response.StatusCode())
	}

	return nil
}
