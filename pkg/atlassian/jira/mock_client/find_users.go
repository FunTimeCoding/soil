package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) FindUsers(_ string) ([]jira.User, error) {
	return nil, nil
}
