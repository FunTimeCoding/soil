package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/team"

func (c *Client) AddTeam(t *team.Team) {
	c.teams = append(c.teams, t)
}
