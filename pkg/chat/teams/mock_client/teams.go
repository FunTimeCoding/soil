package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/team"

func (c *Client) Teams() ([]*team.Team, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return append([]*team.Team{}, c.teams...), nil
}
