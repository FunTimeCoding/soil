package mock_jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) Issue(string) (*issue.Issue, error) { return nil, nil }
