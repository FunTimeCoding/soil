package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"

type Client struct {
	Found    []*client.JiraIssue
	Failure  error
	Searches int
	Closed   int
}
