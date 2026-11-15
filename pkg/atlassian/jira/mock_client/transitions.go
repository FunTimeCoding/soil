package mock_client

import "github.com/andygrunwald/go-jira"

func (c *Client) Transitions(string) ([]jira.Transition, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
