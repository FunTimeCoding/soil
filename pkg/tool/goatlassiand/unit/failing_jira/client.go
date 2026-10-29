package failing_jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"

type Client struct {
	*mock_client.Client
	failure error
}
