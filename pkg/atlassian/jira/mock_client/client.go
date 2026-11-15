package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/failure"
)

type Client struct {
	failures *failure.Failure
	Issues   []*issue.Issue
	Failure  error
	Closed   int
}
