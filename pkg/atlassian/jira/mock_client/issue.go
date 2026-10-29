package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) Issue(string) (*issue.Issue, error) { return nil, nil }
