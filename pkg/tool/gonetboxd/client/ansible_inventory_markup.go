package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
)

func (c *Client) AnsibleInventoryMarkup(field string) (string, error) {
	result, e := c.client.GetAnsibleInventoryWithResponse(
		c.context,
		&client.GetAnsibleInventoryParams{Field: &field},
	)

	if e != nil {
		return "", e
	}

	if result.StatusCode() != http.StatusOK {
		return "", detail_error.New(
			fmt.Sprintf("ansible inventory status: %d", result.StatusCode()),
			result.Status(),
		).WithBody(result.Body)
	}

	return string(result.Body), nil
}
