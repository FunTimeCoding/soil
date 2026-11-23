package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
)

func (c *Client) Teams() ([]*team.Team, error) {
	items, e := c.graph.Items(constant.TeamsPath, nil, 0)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.Team](items)

	if e != nil {
		return nil, e
	}

	result := make([]*team.Team, 0, len(decoded))

	for _, t := range decoded {
		v := team.New(t.Identifier, t.DisplayName)
		v.Description = t.Description
		result = append(result, v)
	}

	return result, nil
}
