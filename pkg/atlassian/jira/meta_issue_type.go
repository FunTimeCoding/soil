package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) MetaIssueType(
	p *jira.MetaProject,
	issueType string,
) (*jira.MetaIssueType, error) {
	result := p.GetIssueTypeWithName(issueType)

	if result == nil {
		return nil, not_found.New("issue type", issueType)
	}

	return result, nil
}
