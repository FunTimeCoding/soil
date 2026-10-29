package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DefaultTeam() *model.Team {
	return &model.Team{Id: "team", Name: "team"}
}
