package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SearchPosts(
	_ *model.Team,
	_ string,
) (*model.PostList, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return model.NewPostList(), nil
}
