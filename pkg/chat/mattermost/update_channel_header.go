package mattermost

import (
	"fmt"
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (c *Client) UpdateChannelHeader(
	h *model.Channel,
	prefix string,
	text string,
	separator string,
) (*model.Channel, error) {
	var header []string
	found := false

	for _, s := range strings.Split(h.Header, separator) {
		if strings.HasPrefix(strings.TrimSpace(s), prefix) {
			header = append(header, fmt.Sprintf(" %s ", text))
			found = true
		} else {
			header = append(header, s)
		}
	}

	if !found {
		header[len(header)-1] = fmt.Sprintf("%s ", header[len(header)-1])
		header = append(header, fmt.Sprintf(" %s", text))
	}

	h.Header = strings.Join(header, separator)
	result, _, e := c.client.UpdateChannel(c.context, h)

	return result, wrapError(e)
}
