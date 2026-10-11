package technitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"net/url"
)

func (c *Client) ZoneOptions(zone string) (*response.ZoneOptions, error) {
	var result response.ZoneOptions

	return &result, c.get(
		fmt.Sprintf(
			"%s?%s",
			constant.ZoneOptionsGetPath,
			url.Values{
				constant.ZoneParameter:              {zone},
				constant.AvailableKeyNamesParameter: {"true"},
			}.Encode(),
		),
		&result,
	)
}
