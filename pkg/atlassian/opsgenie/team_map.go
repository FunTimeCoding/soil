package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"

func (c *Client) TeamMap() *team_map.Map {
	if c.teamMap == nil {
		c.teamMap = team_map.New(c.Teams())
		c.teamMap.TagToName(c.tagToTeam)
	}

	return c.teamMap
}
