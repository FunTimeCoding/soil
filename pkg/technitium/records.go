package technitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium/record"
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"net/url"
)

func (c *Client) Records(
	domain string,
	listZone bool,
) ([]*record.Record, error) {
	var result response.Records
	path := fmt.Sprintf("/zones/records/get?domain=%s", url.QueryEscape(domain))

	if listZone {
		path = fmt.Sprintf("%s&listZone=true", path)
	}

	return result.Records, c.get(path, &result)
}
