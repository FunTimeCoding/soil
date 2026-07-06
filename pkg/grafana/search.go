package grafana

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func (c *Client) Search() models.HitList {
	r, e := c.client.Search.Search(&search.SearchParams{})
	errors.PanicOnError(e)

	return r.Payload
}
