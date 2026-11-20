package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
)

func (o *Tester) Client(j *mock_client.Client) *client.Client {
	o.t.Helper()
	result, _ := o.Reported(j)

	return result
}
