package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) EditMeta(key string) (*jira.EditMetaInfo, error) {
	result, _, e := c.client.Issue.GetEditMeta(&jira.Issue{Key: key})

	return result, wrapError(e)
}
