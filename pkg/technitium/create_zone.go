package technitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"net/url"
)

func (c *Client) CreateZone(
	name string,
	zoneType string,
) (string, error) {
	var result response.CreateZone

	return result.Domain, c.get(
		fmt.Sprintf(
			"/zones/create?zone=%s&type=%s",
			url.QueryEscape(name),
			url.QueryEscape(zoneType),
		),
		&result,
	)
}
