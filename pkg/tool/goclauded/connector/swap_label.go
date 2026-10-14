package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
)

func (c *Client) swapLabel(
	sessionIdentifier string,
	key string,
	expect string,
	value string,
	from string,
) (bool, error) {
	response, e := c.generated.PostSessionLabelWithResponse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionLabelJSONRequestBody{
			Key:    key,
			Value:  &value,
			Expect: &expect,
			From:   &from,
		},
	)

	if e != nil {
		return false, e
	}

	switch response.StatusCode() {
	case http.StatusOK:
		return true, nil
	case http.StatusConflict:
		return false, nil
	default:
		return false, unexpected.Format(
			"label status %d",
			response.StatusCode(),
		)
	}
}
