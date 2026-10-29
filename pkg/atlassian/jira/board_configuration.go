package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) BoardConfiguration(identifier int) (*jira.BoardConfiguration, error) {
	result, r, e := c.client.Board.GetBoardConfiguration(identifier)

	return result, WrapError(r, e)
}
