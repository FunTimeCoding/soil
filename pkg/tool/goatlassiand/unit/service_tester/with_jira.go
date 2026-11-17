package service_tester

import "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"

func WithJira(c *mock_client.Client) Option {
	return func(o *Tester) {
		o.Jira = c
	}
}
