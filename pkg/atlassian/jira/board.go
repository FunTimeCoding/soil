package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Board(identifier int) (*jira.Board, error) {
	result, _, e := c.client.Board.GetBoard(identifier)

	return result, wrapError(e)
}
