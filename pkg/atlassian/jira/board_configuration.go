package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) BoardConfiguration(identifier int) (*jira.BoardConfiguration, error) {
	result, _, e := c.client.Board.GetBoardConfiguration(identifier)

	return result, wrapError(e)
}
