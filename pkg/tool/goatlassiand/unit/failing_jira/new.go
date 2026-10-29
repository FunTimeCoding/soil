package failing_jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"

func New(failure error) *Client {
	return &Client{Client: mock_client.New(), failure: failure}
}
