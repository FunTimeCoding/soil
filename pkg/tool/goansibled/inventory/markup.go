package inventory

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net/http"
	"strings"
)

func (i *Inventory) Markup() (string, error) {
	response, e := i.client.GetAnsibleInventoryWithResponse(
		i.context,
		&client.GetAnsibleInventoryParams{},
	)

	if e != nil {
		return "", e
	}

	if response.StatusCode() != http.StatusOK {
		return "", unexpected.Format(
			"status %d: %s",
			response.StatusCode(),
			strings.TrimSpace(string(response.Body)),
		)
	}

	return string(response.Body), nil
}
