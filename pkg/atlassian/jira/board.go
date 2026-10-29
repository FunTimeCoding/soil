package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Board(identifier int) (*jira.Board, error) {
	result, r, e := c.client.Board.GetBoard(identifier)

	return result, WrapError(r, e)
}
