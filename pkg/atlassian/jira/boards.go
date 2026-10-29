package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Boards() ([]*jira.Board, error) {
	var result []*jira.Board
	v, r, e := c.client.Board.GetAllBoards(nil)

	if e != nil {
		return nil, WrapError(r, e)
	}

	for _, b := range v.Values {
		result = append(result, &b)
	}

	return result, nil
}
