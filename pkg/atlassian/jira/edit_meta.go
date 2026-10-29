package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) EditMeta(key string) (*jira.EditMetaInfo, error) {
	result, r, e := c.client.Issue.GetEditMeta(&jira.Issue{Key: key})

	return result, WrapError(r, e)
}
