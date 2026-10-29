package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) CreateMeta(key string) (*jira.CreateMetaInfo, error) {
	result, r, e := c.client.Issue.GetCreateMeta(key)

	return result, WrapError(r, e)
}
