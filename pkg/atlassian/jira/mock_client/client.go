package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

type Client struct {
	Issues []*issue.Issue
	Closed int
}
